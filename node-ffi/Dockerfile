FROM rust:1.72.0-alpine as builder

WORKDIR /app

RUN apk add --no-cache musl-dev
RUN apk add --update nodejs npm

RUN rustup target add x86_64-unknown-linux-musl

COPY ffi .

RUN npm install
RUN npm run build-release

# RUN cargo build --release

# CMD ["./target/release/rust"]

# FROM node:18-alpine
# WORKDIR /app

# COPY package.json .
# COPY yarn.lock .
# RUN yarn install

# COPY ecosystem.config.js .
# COPY index.mjs .
# RUN yarn global add pm2

# CMD ["pm2-runtime", "ecosystem.config.js"]