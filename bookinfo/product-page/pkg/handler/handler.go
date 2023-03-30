package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/CosmWasm/tinyjson"
	"github.com/product_page/pkg/client"
	"github.com/product_page/pkg/products"
	"github.com/product_page/pkg/template"
)

var detailsTEMP = products.ProductDetails{
	ID:        0,
	Type:      "Paperback",
	Pages:     200,
	Publisher: "PublisherA",
	Language:  "English",
	ISBN10:    "1234567890",
	ISBN13:    "123-1234567890",
}

var reviewsTemp = products.ProductReviews{
	ID:          1,
	PodName:     "reviews-v2-65c4dc6fdc-6bgv9",
	ClusterName: "temp",
	Reviews: []products.Review{
		{
			Reviewer: "Reviewer1",
			Text:     "An extremely entertaining play by Shakespeare. The slapstick humour is refreshing!",
			Rating: products.Rating{
				Stars: 5,
				Color: "",
			},
		},
		{
			Reviewer: "Reviewer2",
			Text:     "Absolutely fun and entertaining. The play lacks thematic depth when compared to other plays by Shakespeare.",
			Rating: products.Rating{
				Stars: 4,
				Color: "green",
			},
		},
	},
}

type Handler struct {
	ProductHandler  *products.ProductHandler
	Client          *client.Client
	template        template.TemplateHandler
	servicesDetails *client.ServicesDetails
}

func NewHandler() *Handler {
	servicesDetails := client.NewServicesDetails()
	return &Handler{
		ProductHandler:  products.NewProductHandler(),
		Client:          client.NewClient(servicesDetails),
		template:        *template.NewTemplateHandler(),
		servicesDetails: servicesDetails,
	}
}

func (h *Handler) ProductPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	productID := 0 // default from istio
	product := h.ProductHandler.GetProduct(productID)
	if product == nil {
		fmt.Println("product not found")
		w.WriteHeader(http.StatusInternalServerError)
	}

	details, status := h.Client.GetDetails(product.ID)
	if status != 200 {
		// w.WriteHeader(status)
		// return
	}
	reviews, status := h.Client.GetReviews(product.ID)
	if status != 200 {
		// w.WriteHeader(status)
		// return
	}
	// TEMP REPLACE
	details = &detailsTEMP
	reviews = &reviewsTemp

	template := h.template.TemplateProductPage(product, details, reviews)
	if _, err := w.Write([]byte(template)); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) ProductsRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("all products requested")
	w.Header().Set("Content-Type", "json")
	products := h.ProductHandler.GetProducts()
	b, err := tinyjson.Marshal(products)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) ProductRoute(w http.ResponseWriter, r *http.Request) {
	// the URL.Query method uses reflection => we cant use this
	id, ok := h.getIdFromUrl(r.URL.Path, 2)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("product %d details requested\n", id)
	w.Header().Set("Content-Type", "json")

	// we would also call the details page in this handler, for now return the
	// base product if it exists
	productDetails, status := h.Client.GetDetails(id)
	if status != 200 {
		w.WriteHeader(status)
		return
	}
	// inefficient to unmarshal then marshal again, but keeps things easier ¯\_(ツ)_/¯
	// todo: new method returning bytes not struct
	b, err := tinyjson.Marshal(productDetails)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) UnimplementedRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("unimplemented route hit")
	w.Header().Set("Content-Type", "json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(string(body))

	fmt.Println(r)
}

func (h *Handler) getIdFromUrl(url string, index int) (int, bool) {
	ss := strings.Split(url, "/")
	if len(ss) <= index {
		return 0, false
	}
	idInt, err := strconv.Atoi(ss[index])
	if err != nil {
		return 0, false
	}
	return idInt, true
}
