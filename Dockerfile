FROM golang:1.21.6-alpine3.19

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build
ADD . .

RUN go build -o bin/main cmd/main.go

ENTRYPOINT ["/build/bin/main"]