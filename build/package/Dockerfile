FROM node:12.14 AS frontbuilder

RUN mkdir /app
WORKDIR /app
ADD . .
RUN cd ./web && yarn install && yarn build


FROM golang:1.13 AS gobuilder

ENV GO111MODULE=on
WORKDIR /go/src/cerberus
COPY --from=frontbuilder /app /go/src/cerberus
RUN go get -u github.com/gobuffalo/packr/v2/packr2
RUN packr2
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/cerberus ./main.go

FROM scratch

COPY --from=gobuilder /go/src/cerberus/bin/cerberus /
EXPOSE 8970
EXPOSE 8971
ENTRYPOINT ["/cerberus", "start"]
