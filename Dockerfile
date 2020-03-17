FROM golang:1.13-alpine

WORKDIR /src/

RUN apk add gcc build-base git mercurial

ADD go.mod .
ADD go.sum .

RUN go mod tidy
