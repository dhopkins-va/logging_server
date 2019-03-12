FROM golang:1.12
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -d -v ./...
RUN go build -o server
CMD ["/app/server"]
