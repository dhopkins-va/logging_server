FROM tinywarrior/logging_server
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o server
CMD ["/app/server"]