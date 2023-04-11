package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	ratings "github.com/ratings_page/pkg/ratings"
)

type Handler struct {
	RatingHandler *ratings.RatingHandler
}

func NewHandler() *Handler {
	return &Handler{
		RatingHandler: ratings.NewRatingHandler(),
	}
}

func (h *Handler) RatingsRoute(w http.ResponseWriter, r *http.Request) {
	log.Println("/ratings: All ratings requested")
	w.Header().Set("Content-Type", "json")
	ratings := h.RatingHandler.GetRatings()
	b, err := json.Marshal(ratings)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) RatingRoute(w http.ResponseWriter, r *http.Request) {
	id, ok := h.getIdFromUrl(r.URL.Path, 2)
	if !ok {
		log.Println("/ratings/{id}: Received invalid rating id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Printf("/rating/%d: Rating for id=%d requested\n", id, id)
	w.Header().Set("Content-Type", "json")
	rating := h.RatingHandler.GetRating(id)

	b, err := json.Marshal(rating)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(b); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) UnimplementedRoute(w http.ResponseWriter, r *http.Request) {
	log.Println("/: Hit default unimplemented route")
	w.Header().Set("Content-Type", "json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(string(body))

	log.Println(r)
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
