FROM golang:1.12

COPY . /go/src/github.com/tinywarrior/logging_server
WORKDIR /go/src/github.com/tinywarrior/logging_server

RUN go get -d -v ./...
RUN go build -o server

CMD ["/go/src/github.com/tinywarrior/logging_server/server"]