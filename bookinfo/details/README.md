Based on: https://github.com/istio/istio/blob/b3d6b939ea66995427585f4bb687ebd7d7cc0741/samples/bookinfo/src/details/details.rb

tinygo version 0.28.0-dev-3104362 darwin/amd64 (using go version go1.19.1 and LLVM version 15.0.0)

spin --version
spin 1.0.0 (df99be2 2023-03-21)

curl http://localhost:3000/health

{"status":"Details is healthy"}

curl http://localhost:3000/details/1

{"id":1,"author":"William Shakespeare","year":"1595","type":"paperback","pages":200,"publisher":"PublisherA","language":"English","ISBN-10":"1234567890","ISBN-13":"123-1234567890"}