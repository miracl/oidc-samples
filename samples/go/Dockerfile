FROM ghcr.io/miracl/oidc-samples/proxy:latest as proxy

FROM golang:alpine as builder
WORKDIR /src
COPY . .
RUN apk add --no-cache git && go build -o /bin/sample .

FROM alpine
COPY --from=proxy /src/ca.crt /usr/local/share/ca-certificates/ca.pem
COPY --from=builder /bin/sample /bin/sample
COPY entrypoint.sh /bin/entrypoint.sh
RUN apk add --no-cache ca-certificates && update-ca-certificates
EXPOSE 8000
ENTRYPOINT entrypoint.sh
