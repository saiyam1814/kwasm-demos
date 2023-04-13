FROM --platform=${BUILDPLATFORM} rust:1.67 AS build
RUN rustup target add wasm32-wasi
COPY . /reviews
WORKDIR /reviews
RUN cargo build --target wasm32-wasi --release
RUN echo "http  80/tcp\nhttp  80/udp" > /etc/services

FROM scratch
COPY --from=build /reviews/target/wasm32-wasi/release/reviews.wasm /target/wasm32-wasi/release/reviews.wasm
COPY ./spin.toml /spin.toml
COPY --from=build /etc/services /etc/services
ENTRYPOINT ["/"]
