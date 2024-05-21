FROM golang:1.20.14-alpine3.18 AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

WORKDIR $GOPATH/src/mypackage/myapp/cmd/server

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/studio-sol-backend-test-2

FROM scratch

COPY --from=builder /go/bin/studio-sol-backend-test-2 /go/bin/studio-sol-backend-test-2

CMD ["/go/bin/studio-sol-backend-test-2"]

EXPOSE 8080