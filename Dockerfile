# Container to compile the app
FROM golang:1.13-alpine AS grpcserver

WORKDIR /build

COPY . .

RUN go build -o /app/grpc-server -mod=vendor
#o - output
# Final container image
FROM alpine:latest

WORKDIR /app

COPY --from=grpcserver /app/grpc-server .

ENTRYPOINT ["/app/grpc-server"]