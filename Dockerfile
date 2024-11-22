FROM golang:1.23.2 AS build

ENV GOARCH=amd64
ENV GOPROXY=https://proxy.golang.org,https://goproxy.io,https://goproxy.dev

RUN apt update -y && apt install -y git make g++ bash curl

WORKDIR /go/src/github.com/service
COPY go.mod ./go.mod
COPY go.sum ./go.sum
COPY cmd ./cmd
COPY internal ./internal
COPY migrations ./migrations

RUN go mod download
RUN go test ./...
RUN mkdir bin
RUN go build  -o bin/main cmd/core/main.go

# REASON I USE 2 STEP DOCKER BUILD
# 1. It is more secure, the docker file is not containig source code
# 2. The docker image is more lightweight
FROM alpine:latest as base
WORKDIR /app
RUN apk add bash nano curl # tzdata  build-base gcompat tcpdump

COPY --from=build /go/src/github.com/service/bin ./bin
COPY --from=build /go/src/github.com/service/migrations /app/migrations

CMD ["/app/bin/main"]

