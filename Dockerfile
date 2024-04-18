FROM golang:latest

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /app
RUN mkdir "/build"
COPY . .

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o /build/app ./cmd" -command="/build/app"

RUN apt-get update
RUN apt-get install sudo
RUN apt-get install nano
