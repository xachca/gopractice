FROM golang:1.9.5-alpine3.6

RUN apk add --no-cache --virtual .build-deps \
            bash \
            git \
            gcc \
            build-base

RUN go get github.com/derekparker/delve/cmd/dlv
EXPOSE 2345
WORKDIR /go/src

CMD ["dlv", "debug", "--listen=:2345", "--headless", "--log", "--"]
