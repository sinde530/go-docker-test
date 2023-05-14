# Build stage
FROM golang:1.18.1-alpine AS build

# Set the environment for the build target
ENV GOOS=linux
ENV GOARCH=arm
ENV GOARM=7

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go ./
RUN CGO_ENABLED=0 go build -tags netgo -ldflags "-extldflags -static -w -s" -o /go-docker-test

# Final stage
FROM scratch

COPY --from=build /go-docker-test /go-docker-test
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

CMD ["/go-docker-test"]
