FROM ubuntu:latest

COPY . /app

CMD ["/app/bin/pixelify"]