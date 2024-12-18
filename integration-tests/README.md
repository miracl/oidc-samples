# MIRACL Trust Integration Tests for OIDC Samples

Integration test suit continuously testing the MIRACL Trust OpenID Connect
implementation with popular client side implementations.

## Getting started

The integration test defines some flags. Use the following command to see them.

### Setup

The integration tests can be run directly with Go or with a Docker container.
You must pass the app's credentials to the tests through environment
variables as follows:

``` bash
export CLIENT_ID=<client-id>
export CLIENT_SECRET=<client-secret>
```

### Run

```bash
go test . \
  -client-id $CLIENT_ID \
  -client-secret $CLIENT_SECRET
```

### Build and run the test binary

```bash
go test -mod=vendor -c -o integration-tests .
./integration-tests \
  -client-id $CLIENT_ID \
  -client-secret $CLIENT_SECRET
```

### See flags

```bash
go test . --args --help
```

### Run in container

```bash
docker run \
  --network host \
  ghcr.io/miracl/oidc-samples/integration-tests:latest \
    --client-id $CLIENT_ID \
    --client-secret $CLIENT_SECRET
```
