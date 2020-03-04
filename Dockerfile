FROM golang:1.14

WORKDIR /go/src/github.com/MrFuku/gasshuku

ADD . .

RUN go get

EXPOSE 8080

RUN cd /go/src/github.com/MrFuku/gasshuku && go build .

CMD ["gasshuku"]
