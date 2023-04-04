# BookInfo Ratings Page

Basic Go implementation of the Istio [Bookinfo sample](https://github.com/istio/istio/tree/master/samples/bookinfo), intended for compilation with the [dev branch](https://github.com/tinygo-org/tinygo/tree/dev) of TinyGo, which has sufficient support for the Go reflect package to not warrant use of TinyJSON.


### The currently servicable endpoints
```
/ratings
/ratings/{product_id}
```

### TO DO:
* Add /health endpoint
* Implement external persistant storage to enable updating book ratings
* Implement toggling of persistant storage via env variables
* Implement special handling for product IDs that are nonexistant