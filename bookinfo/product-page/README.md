# Product Page

tinyGo spin implementation of [istio book-info container](https://github.com/istio/istio/tree/master/samples/bookinfo)


## Endpoints

| Route                              | Docs  |
|------------------------------------|----------------------------------------------------------------------------|
| `/`                                | Displays a templated summary of services and endpoints                     |
| `/productpage`                     | Templated book information, defaults to book ID 1 (currently no override)  |
| `/api/v1/products/{id}`            | Fetch the product details of a given book ID                               |
| `/api/v1/products/{id}/reviews`    | Fetch the product reviews of a given book ID                               |
| `/api/v1/products/{id}/ratings`    | Fetch the product ratings of a given book ID                               |


## JSON marshalling / unmarshalling
These endpoints all use [tinygo](https://github.com/CosmWasm/tinyjson) for json marshalling/unmarshalling. Tinyjson generates custom encoders and decoders bypassing the requirement for reflection, allowing us to implement these handlers in tinygo.
