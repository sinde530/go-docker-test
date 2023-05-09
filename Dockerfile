# syntax=docker/dockerfile:1
# build stage
FROM golang:alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY cmd/main.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-docker-test

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /go-docker-test .

# Options:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
EXPOSE 8080

USER nonroot:nonroot

CMD ["/go-docker-test"]

