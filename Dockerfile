# build stage
FROM golang:1.13-alpine AS build-go
RUN apk --no-cache add git
RUN apk --update add ca-certificates

ENV CODE_PATH /go/src/github.com/velopaymentsapi/payor-example-go

RUN mkdir -p ${CODE_PATH}
COPY . ${CODE_PATH}
WORKDIR ${CODE_PATH}

RUN mkdir -p /tmp/dist
RUN go build cmd/payor/main.go && mv main /tmp/dist/cmd

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-go /tmp/dist/* /app/
COPY --from=build-go /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD ["/app/cmd"]
