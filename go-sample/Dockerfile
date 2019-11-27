FROM golang:alpine as builder
WORKDIR /src
COPY cmd/demo .
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static" -w -s' -o /bin/demo .

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/demo /
COPY --from=builder /src/templates /usr/share/demo/templates
ENTRYPOINT ["/demo", "-templates-dir", "/usr/share/demo/templates"]
