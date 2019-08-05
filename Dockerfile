FROM golang:1.12.6-alpine

WORKDIR /go/goproduct

COPY ./ /go/goproduct

RUN apk add --no-cache \
        alpine-sdk \
        git && \
    go get github.com/labstack/echo && \
    go get github.com/labstack/echo/middleware && \ 
    go get github.com/pilu/fresh

CMD [ "fresh" ]