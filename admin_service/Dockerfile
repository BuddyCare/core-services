FROM golang:1.22-alpine AS builder
ARG BINARY_NAME="admin_service_binary"
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod vendor
RUN gofmt -w .
RUN go clean -modcache
RUN GO111MODULE=on go build -o $BINARY_NAME .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/$BINARY_NAME .
ENTRYPOINT ["/app/$BINARY_NAME"]