FROM golang:1.22-alpine

COPY . /go/src/testy
WORKDIR /go/src/testy

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /go/bin/testy

CMD ["/go/bin/testy"]
