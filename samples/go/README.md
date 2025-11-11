# MIRACL Trust Go OIDC Integration Sample

![go](https://github.com/miracl/oidc-samples/workflows/go/badge.svg)

This is an example of an [OIDC integration](https://miracl.com/resources/docs/guides/authentication/oidc/)
with the [MIRACL Trust platform](https://miracl.com) using Go.
It uses the [go-oidc](https://github.com/coreos/go-oidc) library
to integrate with the MIRACL Trust platform. See its documentation for
the integration steps.

# Setup

To start an OIDC integration, you must create an OIDC application in the
[MIRACL Trust Portal](https://trust.miracl.com) as described [here](https://miracl.com/resources/docs/get-started/low-code/).
The `Redirect URL` must be the same as the one the sample is run with. If you use the
sample's default value, it must be set to `http://localhost:8000/login`.

# Usage

To run the sample, do the following:

```bash
cd samples/go
go run . \
    -client-id <client-id> \
    -client-secret <client-secret> \
    -issuer <oidc-issuer>
```

This starts the sample HTTP server. When you access http://localhost:8000/, you
will be directed to the MIRACL Trust authorization page to begin the
registration and authentication process. After a successful authentication, the
sample retrieves the OIDC UserInfo endpoint and returns the result.

If you would like to run the sample with Docker, follow the instructions in the
main [README.md](../../README.md#run-with-docker) file.
