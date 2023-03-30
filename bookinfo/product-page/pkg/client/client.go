package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/CosmWasm/tinyjson"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/product_page/pkg/products"
)

const (
	defaultServicesDomain  = ""
	defaultDetailsHostname = "details"
	defaultDetailsPort     = "9080"
	defaultRatingsHostname = "ratings"
	defaultRatingsPort     = "9080"
	defaultReviewsHostname = "reviews"
	defaultReviewsPort     = "9080"

	servicesDomainEnvVar  = "SERVICES_DOMAIN"
	detailsHostnameEnvVar = "DETAILS_HOSTNAME"
	detailsPortEnvVar     = "DETAILS_SERVICE_PORT"
	ratingsHostnameEnvVar = "RATINGS_HOSTNAME"
	ratingsPortEnvVar     = "RATINGS_SERVICE_PORT"
	reviewsHostnameEnvVar = "REVIEWS_HOSTNAME"
	reviewsPortEnvVar     = "REVIEWS_SERVICE_PORT"
)

type Client struct {
	client   http.Client
	services *ServicesDetails
}

type ServicesDetails struct {
	productPage endpoint
	details     endpoint
	reviews     endpoint
}

type endpoint struct {
	name     string
	endpoint string
	children []endpoint
}

func NewClient(services *ServicesDetails) *Client {
	return &Client{
		client:   *http.DefaultClient,
		services: services,
	}
}

func (c *Client) GetDetails(id int) (*products.ProductDetails, int) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%d", c.services.details.name, id), bytes.NewBufferString(""))
	if err != nil {
		fmt.Println("error creating request: ", err)
		return nil, http.StatusInternalServerError
	}
	res, err := spinhttp.Send(req)
	if err != nil {
		fmt.Println("spinhttp error: ", err)
		return nil, http.StatusInternalServerError
	}
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode
	}
	if res.Body == nil {
		return nil, http.StatusBadRequest
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response body: ", err)
		return nil, http.StatusInternalServerError
	}
	productDetails := &products.ProductDetails{}
	if err := tinyjson.Unmarshal(b, productDetails); err != nil {
		fmt.Println("unmarshal error: ", err)
		return nil, http.StatusInternalServerError
	}
	return productDetails, http.StatusOK
}

func (c *Client) GetReviews(id int) (*products.ProductReviews, int) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%d", c.services.reviews.name, id), bytes.NewBufferString(""))
	if err != nil {
		fmt.Println("error creating request: ", err)
		return nil, http.StatusInternalServerError
	}
	res, err := spinhttp.Send(req)
	if err != nil {
		fmt.Println("spinhttp error: ", err)
		return nil, http.StatusInternalServerError
	}
	if res.StatusCode != http.StatusOK {
		return nil, res.StatusCode
	}
	if res.Body == nil {
		return nil, http.StatusBadRequest
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response body: ", err)
		return nil, http.StatusInternalServerError
	}
	ProductReviews := &products.ProductReviews{}
	if err := tinyjson.Unmarshal(b, ProductReviews); err != nil {
		fmt.Println("unmarshal error: ", err)
		return nil, http.StatusInternalServerError
	}
	return ProductReviews, http.StatusOK
}

func getEnvVar(key string, defaultVal string) string {
	val, ok := os.LookupEnv(servicesDomainEnvVar)
	if ok {
		return val
	}
	return defaultVal
}

func NewServicesDetails() *ServicesDetails {
	servicesDomain := getEnvVar(servicesDomainEnvVar, defaultServicesDomain)
	detailsHostname := getEnvVar(detailsHostnameEnvVar, defaultDetailsHostname)
	detailsPort := getEnvVar(detailsPortEnvVar, defaultDetailsPort)
	ratingsHostname := getEnvVar(ratingsHostnameEnvVar, defaultRatingsHostname)
	ratingsPort := getEnvVar(ratingsPortEnvVar, defaultRatingsPort)
	reviewsHostname := getEnvVar(reviewsHostnameEnvVar, defaultReviewsHostname)
	reviewsPort := getEnvVar(reviewsPortEnvVar, defaultReviewsPort)

	details := endpoint{
		name:     fmt.Sprintf("http://%s%s:%s", detailsHostname, servicesDomain, detailsPort),
		endpoint: "details",
	}

	ratings := endpoint{
		name:     fmt.Sprintf("http://%s%s:%s", ratingsHostname, servicesDomain, ratingsPort),
		endpoint: "ratings",
	}

	reviews := endpoint{
		name:     fmt.Sprintf("http://%s%s:%s", reviewsHostname, servicesDomain, reviewsPort),
		endpoint: "reviews",
		children: []endpoint{
			ratings,
		},
	}

	productPage := endpoint{
		name:     fmt.Sprintf("http://%s%s:%s", detailsHostname, servicesDomain, detailsPort),
		endpoint: "reviews",
		children: []endpoint{
			details,
			reviews,
		},
	}

	return &ServicesDetails{
		productPage: productPage,
		details:     details,
		reviews:     reviews,
	}
}
