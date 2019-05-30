FROM golang:1.11.2
WORKDIR $GOPATH/src/github.com/booua/dashboard-hub

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/dashboard-hub .

EXPOSE 8080

CMD ["./dashboard-hub"]
