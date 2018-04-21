FROM golang:1.9.5-alpine3.6

RUN apk add --no-cache --virtual .build-deps  git
RUN go get github.com/derekparker/delve/cmd/dlv

COPY . /go/src/begin
WORKDIR /go/src/begin
RUN go build  .

ENTRYPOINT ["dlv", "debug", "./begin", "-l", "0.0.0.0:2345", "--headless=true", "--log=true", "--", "server"]
