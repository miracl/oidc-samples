FROM ghcr.io/miracl/oidc-samples/test-runner:latest
WORKDIR /src
COPY . .
RUN go test -mod=vendor -c -o /usr/local/bin/integration-tests .
ENTRYPOINT ["integration-tests"]
