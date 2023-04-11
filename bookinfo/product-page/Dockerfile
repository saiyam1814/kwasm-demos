FROM --platform=${BUILDPLATFORM} tinygo/tinygo-dev:latest AS build
COPY . /product-page
WORKDIR /product-page
RUN go mod tidy
RUN tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm .

FROM scratch
COPY --from=build /product-page/main.wasm /main.wasm
COPY ./spin.toml /spin.toml
ENTRYPOINT ["/"]
