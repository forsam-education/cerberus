FROM golang:1.13 AS builder

ENV GO111MODULE=on

WORKDIR /go/src/cerberus
ADD . .
RUN go get -u github.com/gobuffalo/packr/v2/packr2
RUN packr2
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/cerberus ./main.go

FROM scratch

COPY --from=builder /go/src/cerberus/bin/cerberus /
EXPOSE 8970
EXPOSE 8971
ENTRYPOINT ["/cerberus", "start"]
