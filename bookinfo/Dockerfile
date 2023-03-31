FROM tinygo/tinygo-dev:latest AS go-base

FROM rust:latest AS rust-base
RUN rustup target add wasm32-wasi

FROM go-base AS details
COPY details/details /src
RUN chmod -R a+rwx /src
WORKDIR /src
RUN go mod tidy
RUN tinygo build -target=wasi -gc=leaking -no-debug -o main.wasm .

FROM rust-base AS reviews
COPY reviews /reviews
WORKDIR /reviews
RUN cargo build --target wasm32-wasi --release

FROM scratch
COPY spin.toml /spin.toml
COPY --from=reviews /reviews/target/wasm32-wasi/release/reviews.wasm /reviews.wasm
COPY --from=details /src/main.wasm /details.wasm
