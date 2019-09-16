FROM golang:1.12-alpine

ENV p $GOPATH/src/github.com/abilioesteves/gowebbuilder

RUN apk add --no-cache git mercurial

RUN mkdir -p ${p}

#ADDED THIS FILE TO USE GO GET
ADD ./main.go ${p}

#ADDED THIS FILES TO USE GO MOD
ADD go.mod go.sum ${p}/

WORKDIR ${p}

RUN go get -v ./...
RUN go mod download 
