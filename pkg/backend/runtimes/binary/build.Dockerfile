FROM golang:1.23-alpine:3.19 AS builder

# Copy the handler into the container
WORKDIR /usr/src/build
COPY functionhandler.go /.

RUN GO111MODULE=off CGO_ENABLED=0 go build -o handler.bin .

FROM alpine:3.19

# Create app directory -> referenced by Dockerfile
WORKDIR /usr/src/app

COPY --from=builder /usr/src/build/handler.bin .

