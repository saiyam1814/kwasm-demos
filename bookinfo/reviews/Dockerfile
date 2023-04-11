FROM --platform=${BUILDPLATFORM} rust:latest AS build
RUN rustup target add wasm32-wasi
COPY . /reviews
WORKDIR /reviews
RUN cargo build --target wasm32-wasi --release

FROM scratch
COPY --from=build /reviews/target/wasm32-wasi/release/reviews.wasm /reviews.wasm
COPY ./spin.toml /spin.toml
ENTRYPOINT ["/"]
