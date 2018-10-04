FROM golang:alpine as builder

RUN apk add --no-cache git
RUN go get -u github.com/gorilla/mux

ENV SRC=/go/src/github.com/aystream/bank-account-exercises/src/
RUN mkdir -p '${SRC}'
ADD . ${SRC}


RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/bank-account-exercises github.com/aystream/bank-account-exercises/src


FROM alpine:3.7

WORKDIR /opt
COPY --from=builder /go/bin/bank-account-exercises /opt/bank-account-exercises

EXPOSE 8080

ENV NAME "Alexander Yakovlev"
ENTRYPOINT /opt/bank-account-exercises