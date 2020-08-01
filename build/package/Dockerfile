FROM golang:1.14 AS gobuilder

ENV GO111MODULE=on
WORKDIR /go/src/cerberus
ARG version_hash

ADD . /go/src/cerberus

RUN go get -u github.com/gobuffalo/packr/v2/packr2

RUN packr2
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'github.com/forsam-education/cerberus/utils.VersionHash=${version_hash}' -X 'github.com/forsam-education/cerberus/utils.BuildTime=`date +%Y-%m-%dT%T%z`'" -o bin/cerberus

FROM scratch

COPY --from=gobuilder /go/src/cerberus/bin/cerberus /
EXPOSE 8970
EXPOSE 8971
ENTRYPOINT ["/cerberus", "start"]
