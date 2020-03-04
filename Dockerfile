FROM golang:1.14

WORKDIR /go/src/github.com/MrFuku/gasshuku

ENV GO111MODULE=on

ADD . .

EXPOSE 8080

RUN go get github.com/pilu/fresh

CMD ["fresh"]
