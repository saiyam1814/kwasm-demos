package handler

import (
	"encoding/json"
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
