FROM golang:1.14

WORKDIR /go/src/app

ADD . .

RUN go get

EXPOSE 8080

CMD ["go", "run", "main.go"]
