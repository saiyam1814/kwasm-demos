package ratings

//tinyjson:skip
type RatingHandler struct{}

func NewRatingHandler() *RatingHandler {
	return &RatingHandler{}
}

//tinyjson:json
type Ratings map[string]int

//tinyjson:json
type Rating struct {
	ID      int     `json:"id"`
	Ratings Ratings `json:"ratings"`
}

func (p *RatingHandler) GetRating(id int) Rating {
	return Rating{
		ID: id,
		Ratings: map[string]int{
			"Reviewer1": 5,
			"Reviewer2": 4,
		},
	}
}
