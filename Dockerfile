FROM golang:1.17-alpine

WORKDIR  ./go/src/git.paygear.ir/cloud/${PROJECT}

RUN go get -v

RUN go build -v -o /go/bin/chemex

COPY . .
COPY chemex /

CMD ["chemex"]