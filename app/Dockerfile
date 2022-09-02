FROM golang:1.17-alpine AS builder

WORKDIR /go/app
COPY ./src /go/app

RUN go mod tidy && go build -o budget-app /go/app/cmd/web



FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder /go/app /app/

EXPOSE 8080

CMD ["./budget-app"]
