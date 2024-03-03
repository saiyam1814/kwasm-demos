ARG BASE_RUN_IMAGE=scratch
FROM docker.io/library/rust:latest AS builder
ARG GITHUB_USER=podtato-head
LABEL org.opencontainers.image.source "https://github.com/${GITHUB_USER}/podtato-head"

RUN rustup target add wasm32-wasi

WORKDIR /build

COPY Cargo.lock .
COPY Cargo.toml .
COPY src/main.rs src/
RUN cargo fetch

COPY . .
RUN cargo build --target wasm32-wasi --release

# must redeclare ARG within build
ARG BASE_RUN_IMAGE
FROM ${BASE_RUN_IMAGE}
ARG GITHUB_USER=podtato-head
LABEL org.opencontainers.image.source "https://github.com/${GITHUB_USER}/podtato-head"

COPY --from=builder /build/target/wasm32-wasi/release/podtato-parts.wasm podtato-parts.wasm

EXPOSE 9000

ENTRYPOINT ["/podtato-parts.wasm"]
