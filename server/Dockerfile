FROM golang:alpine

# Install git & bash
RUN apk update && apk add git && apk add bash

COPY ./src $GOPATH/src/server/src/

COPY src/.env $GOPATH/src/server/.env

WORKDIR  $GOPATH/src/server/src

RUN go get -d -v

WORKDIR  $GOPATH/src/server/