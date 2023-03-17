package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/CosmWasm/tinyjson"
	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/product_page/pkg/products"
)

const (
	detailsPathPrefix = "http://localhost:3001/details/"
)

type Client struct {
	client http.Client
}

func NewClient() *Client {
	return &Client{
		client: *http.DefaultClient,
	}
}

func (c *Client) GetDetails(id int) (*products.ProductDetails, int) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%d", detailsPathPrefix, id), bytes.NewBufferString(""))
	req.Header.Add("foo", "bar")
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
