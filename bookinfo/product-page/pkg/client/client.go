package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CosmWasm/tinyjson"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/product_page/pkg/config"
	"github.com/product_page/pkg/products"
)

type Client struct {
	client   http.Client
	services *config.ServicesConfig
}

func NewClient(services *config.ServicesConfig) *Client {
	return &Client{
		client:   *http.DefaultClient,
		services: services,
	}
}

func (c *Client) GetDetails(id int) (*products.ProductDetails, int) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/details/%d", c.services.Details.Name, id), bytes.NewBufferString(""))
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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/reviews/%d", c.services.Reviews.Name, id), bytes.NewBufferString(""))
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
	productReviews := &products.ProductReviews{}
	if err := tinyjson.Unmarshal(b, productReviews); err != nil {
		fmt.Println("unmarshal error: ", err)
		return nil, http.StatusInternalServerError
	}
	return productReviews, http.StatusOK
}

func (c *Client) GetRatings(id int) (*products.ProductRatings, int) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/ratings/%d", c.services.Ratings.Name, id), bytes.NewBufferString(""))
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
	productRatings := &products.ProductRatings{}
	if err := tinyjson.Unmarshal(b, productRatings); err != nil {
		fmt.Println("unmarshal error: ", err)
		return nil, http.StatusInternalServerError
	}
	return productRatings, http.StatusOK
}
