package main

import (
	"fmt"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go-chi/chi/v5"
)

func main() {}

func init() {
	spinhttp.Handle(func(res http.ResponseWriter, req *http.Request) {
		router := chi.NewRouter()

		router.Mount("/chi/", GetRoutes())

		router.ServeHTTP(res, req)
	})
}

func GetRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", Index)
	router.Get("/hello/{id}", Hello)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s!\n", chi.URLParam(r, "id"))
}
