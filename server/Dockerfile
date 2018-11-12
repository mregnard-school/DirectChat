FROM golang:1.11.2-alpine

# Install git & bash
RUN apk update && apk add git bash gcc dpkg musl-dev

COPY ./src $GOPATH/src/server

WORKDIR  $GOPATH/src/server

RUN go get -d -v