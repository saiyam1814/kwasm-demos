package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/CosmWasm/tinyjson"
	"github.com/product_page/pkg/products"
)

type Handler struct {
	ProductHandler *products.ProductHandler
}

func NewHandler() *Handler {
	return &Handler{
		ProductHandler: products.NewProductHandler(),
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
	fmt.Printf("product %d requested\n", id)
	w.Header().Set("Content-Type", "json")

	// we would also call the details page in this handler, for now return the
	// base product if it exists
	product := h.ProductHandler.GetProduct(id)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	b, err := tinyjson.Marshal(product)
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
