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
				DescriptionHtml: `<a href="https://en.wikipedia.org/wiki/The_Comedy_of_Errors">Wikipedia Summary</a>: The Comedy of Errors is one of <b>William Shakespeare\'s</b> early plays. It is his shortest and one of his most farcical comedies, with a major part of the humour coming from slapstick and mistaken identity, in addition to puns and word play.`,
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
