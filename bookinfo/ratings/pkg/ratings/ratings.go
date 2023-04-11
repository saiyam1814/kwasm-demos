package ratings

import "sync"

//tinyjson:skip
type RatingHandler struct {
	mu      sync.Mutex
	ratings Ratings
}

func NewRatingHandler() *RatingHandler {
	return &RatingHandler{
		ratings: Ratings{
			{
				ID:          0,
				Title:       "The Comedy of Errors",
				Rating:      4,
				RatingCount: 12,
			},
			{
				ID:          1,
				Title:       "Designing Data Intensive Applications",
				Rating:      5,
				RatingCount: 4,
			},
			{
				ID:          2,
				Title:       "WebAssembly for dummies",
				Rating:      4,
				RatingCount: 6,
			},
		},
	}
}

//tinyjson:json
type Ratings []Rating

//tinyjson:json
type Rating struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Rating      float64 `json:"rating"`
	RatingCount int     `json:"ratingcount"`
}

func (p *RatingHandler) GetRatings() Ratings {
	return p.ratings
}

func (p *RatingHandler) GetRating(id int) *Rating {
	for _, rating := range p.ratings {
		if rating.ID == id {
			return &rating
		}
	}
	return nil
}
