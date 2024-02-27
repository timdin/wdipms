FROM golang:latest

RUN mkdir -p /usr/local/go/src/wdipms
WORKDIR /usr/local/go/src/wdipms
ADD . /usr/local/go/src/wdipms

RUN go mod download
RUN go build ./main.go

EXPOSE 8080

CMD ["./main"]