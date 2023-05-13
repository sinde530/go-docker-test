# syntax=docker/dockerfile:1
# Build stage
FROM golang:1.18.1-buster AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o /go-docker-test

# Final stage
FROM --platform=linux/arm/v7 arm32v7/debian:buster-slim

COPY --from=build /go-docker-test /go-docker-test
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["/go-docker-test"]
