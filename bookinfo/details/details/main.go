package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go-chi/chi/v5"
)

type BookDetails struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Year      int    `json:"year"`
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

			details, err := getBookDetails(id)
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

func getBookDetails(id int) (BookDetails, error) {
	if os.Getenv("ENABLE_EXTERNAL_BOOK_SERVICE") == "true" {
		log.Println("Fetching details from external service")
		isbn := "0486424618"
		return fetchDetailsFromExternalService(isbn, id)
	}

	return BookDetails{
		ID:        id,
		Author:    "William Shakespeare",
		Year:      1595,
		Type:      "paperback",
		Pages:     200,
		Publisher: "PublisherA",
		Language:  "English",
		ISBN10:    "1234567890",
		ISBN13:    "123-1234567890",
	}, nil
}

type GoogleBooksResponse struct {
	Items []struct {
		VolumeInfo struct {
			Authors             []string `json:"authors"`
			PublishedDate       string   `json:"publishedDate"`
			PrintType           string   `json:"printType"`
			PageCount           int      `json:"pageCount"`
			Publisher           string   `json:"publisher"`
			Language            string   `json:"language"`
			IndustryIdentifiers []struct {
				Type       string `json:"type"`
				Identifier string `json:"identifier"`
			} `json:"industryIdentifiers"`
		} `json:"volumeInfo"`
	} `json:"items"`
}

func fetchDetailsFromExternalService(isbn string, id int) (BookDetails, error) {
	url := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=isbn:%s", isbn)

	resp, err := spinhttp.Get(url)

	if err != nil {
		return BookDetails{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BookDetails{}, err
	}

	var jsonResponse GoogleBooksResponse
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return BookDetails{}, err
	}

	book := jsonResponse.Items[0].VolumeInfo

	published, err := time.Parse("2006-01-02", book.PublishedDate)
	if err != nil {
		return BookDetails{}, err
	}

	return BookDetails{
		ID:        id,
		Author:    book.Authors[0],
		Year:      published.Year(),
		Type:      book.PrintType,
		Pages:     book.PageCount,
		Publisher: book.Publisher,
		Language:  book.Language,
		ISBN10:    book.IndustryIdentifiers[1].Identifier,
		ISBN13:    book.IndustryIdentifiers[0].Identifier,
	}, nil
}

func main() {}
