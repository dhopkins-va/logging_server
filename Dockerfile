FROM tinywarrior/alpine_go
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o server
CMD ["/app/server"]