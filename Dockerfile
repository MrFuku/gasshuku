FROM golang:1.14

WORKDIR /go/src/app

ADD . .

RUN go get

