FROM golang:1.17

ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip

ENV ZONEINFO /zoneinfo.zip

RUN mkdir /garden

COPY . /garden

WORKDIR /garden

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -o /build

CMD [ "/build" ]