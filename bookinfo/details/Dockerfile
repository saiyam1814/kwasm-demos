FROM --platform=${BUILDPLATFORM} tinygo/tinygo-dev:latest AS build
COPY ./details /details
WORKDIR /details
RUN go mod tidy
RUN tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm .

FROM scratch
COPY --from=build /details/main.wasm /details.wasm
COPY ./details/spin.toml /spin.toml
