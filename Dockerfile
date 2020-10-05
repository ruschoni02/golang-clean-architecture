FROM golang:alpine

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apk add -U --no-cache build-base ca-certificates git make

WORKDIR /app

ADD go.mod go.sum ./

ADD . .

RUN make deps

CMD make dev
