FROM golang:alpine
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/booua/dashboard-hub

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=arm64 go build

EXPOSE 8081

CMD ["./dashboard-hub"]
