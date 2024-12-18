# MIRACL Trust DotNet 6.0 OIDC Integration Sample

![dotnet6](https://github.com/miracl/oidc-samples/workflows/dotnet6/badge.svg)

This is an example of an [OIDC integration](https://miracl.com/resources/docs/guides/authentication/oidc/)
with the [MIRACL Trust platform](https://miracl.com) using DotNet6.0.

# Dependencies

This sample uses [.NET authentication](https://www.nuget.org/packages/Microsoft.AspNetCore.Authentication.OpenIdConnect/)
library to integrate with the MIRACL Trust platform.

# Setup

The sample can be run directly with Go or with a Docker container. If you choose
the latter, follow the instructions [here](../../README.md).

To start an OIDC integration, you must create an OIDC application in the
[MIRACL Trust Portal](https://trust.miracl.com) as described [here](https://miracl.com/resources/docs/get-started/low-code/).
The `Redirect URL` must be the same as the one the sample is run with. If you use the
sample's default value, it must be set to `http://localhost:8000/login`.
You must pass the app's credentials to the sample through environment
variables as follows:

``` bash
export CLIENT_ID=<client-id>
export CLIENT_SECRET=<client-secret>
```

# Usage

To run the sample, do the following:

```bash
cd samples/dotnet6
dotnet run
```

This runs a sample HTTP server and when you open http://localhost:8000/,
it navigates you to the MIRACL Trust authorization page to start the
registration and authentication.
