package main

import (
	"net/http"

	"github.com/ratings_page/pkg/handler"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go-chi/chi/v5"
)

func init() {
	spinhttp.Handle(func(res http.ResponseWriter, req *http.Request) {
		handler := handler.NewHandler()

		router := chi.NewRouter()
		router.Get("/", handler.UnimplementedRoute)
		router.Get("/ratings", handler.RatingsRoute)
		router.Get("/ratings/{id}", handler.RatingRoute)

		router.ServeHTTP(res, req)
	})
}

func main() {}
