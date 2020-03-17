# builder step
FROM abilioesteves/gowebbuilder:1.4.0 as builder

WORKDIR /src/

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /example main.go

# cleaning step
FROM scratch

COPY --from=builder /example /

CMD ["/example"]