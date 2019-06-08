FROM golang:alpine
RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/github.com/booua/dashboard-hub

COPY . .

RUN go get -d -v

RUN GOOS=linux GOARCH=arm GOARM=5 go build -a -ldflags="-w -s" -o /build/dashboard-hub

FROM scratch

COPY --from=builder /build/dashboard-hub /build/dashboard-hub

EXPOSE 8081

CMD ["/build/dashboard-hub"]
