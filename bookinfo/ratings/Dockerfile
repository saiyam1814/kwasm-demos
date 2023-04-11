FROM --platform=${BUILDPLATFORM} ghcr.io/tinygo-org/tinygo-dev:latest AS build
COPY . /ratings
WORKDIR /ratings
RUN go mod tidy
RUN tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm .

FROM scratch
COPY --from=build /ratings/main.wasm /main.wasm
COPY ./spin.toml /spin.toml
ENTRYPOINT ["/"]
