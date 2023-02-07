FROM golang:1.20.0-alpine3.17 AS builder

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www

ENV CGO_ENABLED=0

ENV GOARCH=amd64

ENV GOOS=linux

RUN go build -o app

EXPOSE 3200

ENTRYPOINT ["/var/www/app"]
