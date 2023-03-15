package main

import (
	"fmt"
	"net/http"

	"github.com/suard/pkg/env"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/go-chi/chi/v5"
)

func main() {}

func init() {
	spinhttp.Handle(func(res http.ResponseWriter, req *http.Request) {
		router := chi.NewRouter()

		router.Mount("/", GetRoutes())

		router.ServeHTTP(res, req)
	})
}

func GetRoutes() http.Handler {
	router := chi.NewRouter()
	env := env.New()

	router.Get("/", Index)
	router.Get("/hello/{id}", Hello)
	env.AddRoutes(router)
	// router.Get("/env/api", GetHeaders)

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, %s!\n", chi.URLParam(r, "id"))
}
