FROM rust:1.72.0-alpine

WORKDIR /app

RUN apk add --no-cache musl-dev

COPY . .

RUN cargo build --release

CMD ["./target/release/rust"]
