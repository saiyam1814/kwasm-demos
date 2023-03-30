package template

import (
	"fmt"
	"strings"

	"github.com/product_page/pkg/products"
)

type TemplateHandler struct{}

func NewTemplateHandler() *TemplateHandler {
	return &TemplateHandler{}
}

func (t *TemplateHandler) TemplateProductPage(product products.Product, details products.ProductDetails, reviews products.ProductReviews) string {
	productPage := productPageHTML
	s := t.newSummary(product)
	d := t.newDetails(details)
	r := t.newReviews(reviews)
	productPage = strings.ReplaceAll(productPage, productPageSummaryReplaceTarget, s)
	productPage = strings.ReplaceAll(productPage, productPageDetailsReplaceTarget, d)
	productPage = strings.ReplaceAll(productPage, productPageReviewsReplaceTarget, r)
	return productPage
}

func (t *TemplateHandler) newDetails(productDetails products.ProductDetails) string {
	details := detailsHTML
	details = strings.ReplaceAll(details, detailsTypeReplaceTarget, productDetails.Type)
	details = strings.ReplaceAll(details, detailsPagesReplaceTarget, fmt.Sprintf("%d", productDetails.Pages))
	details = strings.ReplaceAll(details, detailsPublisherReplaceTarget, productDetails.Publisher)
	details = strings.ReplaceAll(details, detailsLanguageReplaceTarget, productDetails.Language)
	details = strings.ReplaceAll(details, detailsISBN10ReplaceTarget, productDetails.ISBN10)
	details = strings.ReplaceAll(details, detailsISBN13ReplaceTarget, productDetails.ISBN13)
	return details
}

func (t *TemplateHandler) newSummary(product products.Product) string {
	summary := summaryHTML
	summary = strings.ReplaceAll(summary, titleReplaceTarget, product.Title)
	summary = strings.ReplaceAll(summary, descriptionReplaceTarget, product.DescriptionHtml)
	return summary
}

func (t *TemplateHandler) newReviews(reviews products.ProductReviews) string {
	reviewsOut := reviewsHTML
	revs := ""
	for _, rev := range reviews.Reviews {
		temp := reviewHTML
		temp = strings.ReplaceAll(temp, reviewsTextReplaceTarget, rev.Text)
		temp = strings.ReplaceAll(temp, reviewsReviewerReplaceTarget, rev.Reviewer)
		stars := ""
		for i := 1; i <= 5; i++ {
			if i <= rev.Rating.Stars {
				stars = stars + filledStar
			} else {
				stars = stars + emptyStar
			}
		}
		temp = strings.ReplaceAll(temp, reviewsStarsReplaceTarget, stars)
		revs = revs + temp
	}
	reviewsOut = strings.ReplaceAll(reviewsOut, reviewsReviewsReplaceTarget, revs)
	reviewsOut = strings.ReplaceAll(reviewsOut, reviewsServedByReplaceTarget, reviews.PodName)
	return reviewsOut
}
