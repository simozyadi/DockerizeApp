FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY main.go .

RUN go build -o webserver .

FROM alpine:3.13

COPY --from=builder /go/src/app/webserver /go/src/app/

WORKDIR /go/src/app/

CMD ["./webserver"]
