FROM golang:1.12-alpine

WORKDIR /src/

RUN apk add --no-cache gcc build-base git mercurial

ADD go.mod .
ADD go.sum .

RUN go mod tidy
