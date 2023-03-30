package main

import (
	"net/http"

	"github.com/product_page/pkg/handler"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go-chi/chi/v5"
)

func init() {
	spinhttp.Handle(func(res http.ResponseWriter, req *http.Request) {
		handler := handler.NewHandler()
		// we need to setup the router inside spin handler
		router := chi.NewRouter()
		router.Get("/", handler.Index)
		router.Get("/productpage", handler.ProductPage)
		router.Get("/api/v1/products", handler.Products)
		router.Get("/api/v1//products/{id}", handler.Product)
		router.Get("/api/v1//products/{id}/reviews", handler.Reviews)
		router.Get("/api/v1//products/{id}/ratings", handler.Ratings)

		// hand the request/response off to chi
		router.ServeHTTP(res, req)
	})
}

func main() {}
