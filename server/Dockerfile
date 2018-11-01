FROM golang:alpine

# Install git & bash
RUN apk update && apk add git && apk add bash

COPY ./src $GOPATH/src/server

WORKDIR  $GOPATH/src/server

RUN go get -d -v