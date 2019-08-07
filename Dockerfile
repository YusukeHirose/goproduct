FROM golang:1.12.6-alpine

WORKDIR /go/src/goproduct

COPY ./ /go/src/goproduct

RUN apk add --no-cache \
        alpine-sdk \
        git && \
    go get github.com/labstack/echo && \
    go get github.com/labstack/echo/middleware && \ 
    go get github.com/jinzhu/gorm && \
    go get github.com/go-sql-driver/mysql && \
    go get github.com/pilu/fresh

CMD [ "fresh" ]