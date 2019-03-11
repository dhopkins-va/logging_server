FROM tinywarrior/alpine_go:v0.1
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o server
CMD ["/app/server"]