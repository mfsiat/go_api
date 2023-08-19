FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY env-example .env
COPY . .
RUN go mod tidy
