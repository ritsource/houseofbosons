FROM golang:alpine

RUN apk update && apk add --no-cache git

ENV PORT 8080

WORKDIR /go/src/github.com/ritsource/houseofbosons/server

COPY . .

RUN go get -d -v

RUN go build -o /go/bin/houseofbosons-server

CMD [ "/go/bin/houseofbosons-server" ]