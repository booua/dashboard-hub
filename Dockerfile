FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN mkdir google

ENV GOOGLE_APPLICATION_CREDENTIALS="/google/dashboard-hub-0dea1d71f6ab.json"

WORKDIR $GOPATH/src/github.com/booua/dashboard-hub

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=arm64 go build

EXPOSE 8080

CMD ["./dashboard-hub"]
