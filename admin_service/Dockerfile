FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod vendor
RUN gofmt -w .
RUN go clean -modcache
RUN GO111MODULE=on go build -o admin_service_binary .

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/admin_service_binary .
ENTRYPOINT ["/app/admin_service_binary"]