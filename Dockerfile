# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /myapp

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./proxy-server ./proxy-server

RUN go build -o /app

EXPOSE 8080

CMD ["/app"]