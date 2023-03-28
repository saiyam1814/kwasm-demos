package main

import (
	"encoding/json"
	"fmt"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type BookDetails struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Year      string `json:"year"`
	Type      string `json:"type"`
	Pages     int    `json:"pages"`
	Publisher string `json:"publisher"`
	Language  string `json:"language"`
	ISBN10    string `json:"ISBN-10"`
	ISBN13    string `json:"ISBN-13"`
}

func init() {
	spinhttp.Handle(func(res http.ResponseWriter, req *http.Request) {
		r := chi.NewRouter()

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"status": "Details is healthy"})
		})

		r.Get("/details/{id}", func(w http.ResponseWriter, r *http.Request) {
			idStr := chi.URLParam(r, "id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "please provide numeric product id", http.StatusBadRequest)
				return
			}

			details, err := getBookDetails(id, nil)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(details)
		})

		r.ServeHTTP(res, req)
	})
}

func getBookDetails(id int, headers map[string]string) (BookDetails, error) {
	if os.Getenv("ENABLE_EXTERNAL_BOOK_SERVICE") == "true" {
		isbn := "0486424618"
		return fetchDetailsFromExternalService(isbn, id, headers)
	}

	return BookDetails{
		ID:        id,
		Author:    "William Shakespeare",
		Year:      "1595",
		Type:      "paperback",
		Pages:     200,
		Publisher: "PublisherA",
		Language:  "English",
		ISBN10:    "1234567890",
		ISBN13:    "123-1234567890",
	}, nil
}

func fetchDetailsFromExternalService(isbn string, id int, headers map[string]string) (BookDetails, error) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s", isbn)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return BookDetails{}, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return BookDetails{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BookDetails{}, err
	}

	var jsonResponse map[string]interface{}
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return BookDetails{}, err
	}

	book := jsonResponse["items"].([]interface{})[0].(map[string]interface{})["volumeInfo"].(map[string]interface{})

	language := "unknown"
	if book["language"] == "en" {
		language = "English"
	}

	bookType := "unknown"
	if book["printType"] == "BOOK" {
		bookType = "paperback"
	}

	isbn10 := getISBN(book, "ISBN_10")
	isbn13 := getISBN(book, "ISBN_13")

	return BookDetails{
		ID:        id,
		Author:    book["authors"].([]interface{})[0].(string),
		Year:      book["publishedDate"].(string),
		Type:      bookType,
		Pages:     int(book["pageCount"].(float64)),
		Publisher: book["publisher"].(string),
		Language:  language,
		ISBN10:    isbn10,
		ISBN13:    isbn13,
	}, nil
}

func getISBN(book map[string]interface{}, isbnType string) string {
	identifiers := book["industryIdentifiers"].([]interface{})
	for _, identifier := range identifiers {
		idMap := identifier.(map[string]interface{})
		if idMap["type"] == isbnType {
			return idMap["identifier"].(string)
		}
	}
	return ""
}

func main() {}
