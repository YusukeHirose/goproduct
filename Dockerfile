FROM golang:1.12.6-alpine

WORKDIR /go/goproduct

COPY ./ /go/goproduct

ENV GO111MODULE=on

RUN apk add --no-cache \
        alpine-sdk \
        git && \
        go get github.com/pilu/fresh
    
CMD [ "fresh" ]