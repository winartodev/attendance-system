FROM golang:1.19.4-alpine AS builder

RUN apk update && apk upgrade && \
    apk --update add git make curl

RUN mkdir /attendance-system

ADD . /attendance-system

WORKDIR /attendance-system

RUN go mod download

RUN go build app/main.go

EXPOSE 8080

ENTRYPOINT ["/attendance-system/main"]