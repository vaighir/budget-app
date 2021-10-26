FROM golang:1.17-alpine

WORKDIR app

COPY ./src /go/app

RUN go mod download

EXPOSE 8080

CMD go run cmd/web/*go
