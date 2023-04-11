FROM --platform=${BUILDPLATFORM} rust:1.67 AS build
RUN rustup target add wasm32-wasi
COPY . /reviews
WORKDIR /reviews
RUN cargo build --target wasm32-wasi --release

FROM scratch
COPY --from=build /reviews/target/wasm32-wasi/release/reviews.wasm /target/wasm32-wasi/release/reviews.wasm
COPY ./spin.toml /spin.toml
ENTRYPOINT ["/"]
