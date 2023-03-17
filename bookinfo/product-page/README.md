# Product Page

Initial implementation fo the product page contains 2 routes:

```
/products
/products/{product_id}
```

These endpoints all use [tinygo](https://github.com/CosmWasm/tinyjson) for json marshalling/unmarshalling. Tinyjson generates custom encoders and decoders bypassing the requirement for reflection, allowing us to implement these handlers in tinygo.

next steps:

include requests for details / reviews for products + extend the handlers to include the remaining endpoints 
HTML templating