FROM alpine:latest
FROM golang:latest

WORKDIR /usr/src/app

COPY . .

RUN go mod download

COPY . .

EXPOSE 7777

CMD go run main.go

# https://www.youtube.com/watch?v=ofdSnD0b1YU
