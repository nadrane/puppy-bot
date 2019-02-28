FROM golang:1.12.0-alpine3.9

WORKDIR /usr/src/app
EXPOSE 8080

COPY puppy.go puppy.go
COPY server.go server.go

RUN go build

CMD ./app