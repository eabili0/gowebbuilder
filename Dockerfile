FROM golang:1.19-alpine

WORKDIR /src/

RUN apk add gcc build-base git mercurial

ADD go.mod .
ADD go.sum .

RUN go mod download
