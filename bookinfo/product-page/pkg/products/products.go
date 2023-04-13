package products

//tinyjson:skip
type ProductHandler struct {
	products Products
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		products: Products{
			{
				ID:              0,
				Title:           "The Comedy of Errors",
				DescriptionHtml: `<a href="https://en.wikipedia.org/wiki/The_Comedy_of_Errors">Wikipedia Summary</a>: The Comedy of Errors is one of <b>William Shakespeare's</b> early plays. It is his shortest and one of his most farcical comedies, with a major part of the humour coming from slapstick and mistaken identity, in addition to puns and word play.`,
			},
		},
	}
}

//tinyjson:json
type Products []Product

//tinyjson:json
type Product struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	DescriptionHtml string `json:"descriptionHtml"`
}

//tinyjson:json
type ProductDetails struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Year      uint16 `json:"year"`
	Type      string `json:"type"`
	Pages     uint16 `json:"pages"`
	Publisher string `json:"publisher"`
	Language  string `json:"language"`
	ISBN10    string `json:"ISBN-10"`
	ISBN13    string `json:"ISBN-13"`
}

//tinyjson:json
type ProductReviews struct {
	ID          string   `json:"id"`
	PodName     string   `json:"podname"`
	ClusterName string   `json:"clustername"`
	Reviews     []Review `json:"reviews"`
}

//tinyjson:json
type Review struct {
	Reviewer string `json:"reviewer"`
	Text     string `json:"text"`
	Rating   Rating `json:"rating"`
}

//tinyjson:json
type Rating struct {
	Stars *int    `json:"stars"`
	Color *string `json:"color"`
	Error *string `json:"error"`
}

//tinyjson:json
type ProductRatings struct {
	ID      int            `json:"id"`
	Ratings map[string]int `json:"ratings"`
}

func (p *ProductHandler) GetProducts() Products {
	return p.products
}

func (p *ProductHandler) GetProduct(id int) *Product {
	for _, product := range p.products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
