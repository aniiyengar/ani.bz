
FROM golang:1.13.5-alpine3.10
WORKDIR /go/src/github.com/aniiyengar/ani.bz
COPY . .

RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go build main.go

CMD ["./main"]
EXPOSE 9003
