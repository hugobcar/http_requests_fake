FROM golang:1.16.3 as builder
WORKDIR /go/src/github.com/hugobcar/http_requests_fake
ADD . /go/src/github.com/hugobcar/http_requests_fake
# RUN GO111MODULE=on go mod vendor
RUN CGO_ENABLED=0 go build -o http_requests_fake

FROM alpine:3.10.1
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/* && apk add bc
WORKDIR /app
COPY --from=builder /go/src/github.com/hugobcar/http_requests_fake/http_requests_fake .
CMD ["./http_requests_fake"]
