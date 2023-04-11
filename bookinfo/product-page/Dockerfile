FROM --platform=${BUILDPLATFORM} tinygo/tinygo-dev:latest AS build
COPY . /details
WORKDIR /details
RUN go mod tidy
RUN tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm .

FROM scratch
COPY --from=build /details/main.wasm /main.wasm
COPY ./details/spin.toml /spin.toml
