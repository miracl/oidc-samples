# MIRACL Trust OIDC Integration Samples

![go](https://github.com/miracl/oidc-samples/workflows/go/badge.svg)
![nodejs](https://github.com/miracl/oidc-samples/workflows/nodejs/badge.svg)
![dotnet6](https://github.com/miracl/oidc-samples/workflows/dotnet6/badge.svg)
![dotnet](https://github.com/miracl/oidc-samples/workflows/dotnet/badge.svg)
![python](https://github.com/miracl/oidc-samples/workflows/python/badge.svg)

This repository contains a collection of samples which integrate between the
[MIRACL Trust platform](https://miracl.com) and various OIDC libraries. This
integration test suite validates both the compatibility and the correct operation
of the client library.

## ENV Variables

All samples work with the following environment variables:

- `HOST` - Host to listen on. The default is "localhost".
- `PORT` - Port of the listening host. The default is "8000".
- `ISSUER` - OpenID Connect Issuer. The default is "https://api.mpin.io".
- `REDIRECT_URL` - The redirect URL of the app in the MIRACL Trust platform.
  The default value is "http://localhost:8000/login".
- `CLIENT_ID` - The client ID of the app in the MIRACL Trust platform. It has
  no default value and is mandatory.
- `CLIENT_SECRET`- The client secret of the app in the MIRACL Trust platform.
  It has no default value and is mandatory.
- `PROXY_HOST`- The host address of the proxy, which you wish to run the sample
  behind. The default value is an empty string.
- `PROXY_PORT`- The port of the proxy, which you wish to run the sample behind.
  The default value is an empty string.

The required environment variables are `CLIENT_ID` and `CLIENT_SECRET`.

To get those values, you'll need to [register](https://miracl.com/resources/docs/get-started/register/) and [create an app](https://miracl.com/resources/docs/get-started/low-code/) in our platform.

## Usage

Every sample can be run in its own framework or as a Docker container as shown
here:

```bash
cd samples/<variant>
docker build -t sample .
docker run \
  --publish 8000:8000 \
  --env CLIENT_ID=<client-id> \
  --env CLIENT_SECRET=<client-secret> \
  sample
```

Then, go to `http://localhost:8000` where you are navigated to the MIRACL Trust
authorization page to start your registration and authentication. If you need
to set another `PORT` or `HOST`, change the `8000:8000` port mapping and add
the environment variables in the `docker run` command.

See the respective README files for details of every sample if you don't want to
use Docker.

## Running through proxy

To test how OIDC libraries behave in some edge cases (e.g. when the
OIDC server misbehaves), we need to modify the traffic between the library and
the sample showcasing it.

You have the option to use our proxy with the provided samples. You can check
the [README](proxy/README.md) in the proxy directory for information on how to
build and run it.

Provided that you have built Docker images of the proxy and the sample that you
wish to run, you can run both `docker run` commands with the addition of the
`PROXY_HOST` and `PROXY_PORT` environment variables. If you use the default values,
the commands to run the sample behind the proxy are:

```bash
docker run \
  --publish 8080:8080 \
  proxy
docker run \
  --publish 8000:8000 \
  --env PROXY_HOST=127.0.0.1 \
  --env PROXY_PORT=8080 \
  --env CLIENT_ID=<client-id> \
  --env CLIENT_SECRET=<client-secret> \
  sample
```

You can confirm that the requests from the sample are going through a proxy if
you enable the verbose mode of the proxy using the `VERBOSE` environment
variable in the command above. When the proxy and the sample are started and you
complete a registration and authentication, the proxy output will log out the
information of the proxied requests.
