# builder step
FROM abilioesteves/gowebbuilder:unstable as builder

ENV p $GOPATH/src/github.com/abilioesteves/gowebbuilder/example/

RUN mkdir -p ${p}
ADD ./main.go ${p}
WORKDIR ${p}

# NOTICE THE LACK OF A `go get` CALL

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /example main.go

# cleaning step
FROM scratch

COPY --from=builder /example /

CMD ["/example"]