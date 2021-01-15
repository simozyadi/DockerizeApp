FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY main.go .

RUN go build -o webserver .


FROM alpine:3.13

WORKDIR /app

COPY --from=builder /go/src/app/webserver /app/

CMD ["./webserver"]
