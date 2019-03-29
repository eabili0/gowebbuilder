FROM golang:1.11-alpine

RUN apk add --no-cache git mercurial

ENV p $GOPATH/src/github.com/abilioesteves/gowebbuilder

RUN mkdir -p ${p}

ADD ./main.go ${p}

WORKDIR ${p}

RUN go get -v ./...