FROM golang:1.21.4-alpine3.17 AS builder

WORKDIR $GOPATH/src/http/
COPY go/main.go go.mod go.sum ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/config-api

FROM scratch
COPY --from=builder /go/bin/config-api /go/bin/bin/config-api

CMD ["/go/bin/bin/config-api"]