# MIRACL Trust OIDC Integration Samples

![go](https://github.com/miracl/oidc-samples/workflows/go/badge.svg)
![nodejs](https://github.com/miracl/oidc-samples/workflows/nodejs/badge.svg)
![dotnet6](https://github.com/miracl/oidc-samples/workflows/dotnet6/badge.svg)
![dotnet](https://github.com/miracl/oidc-samples/workflows/dotnet/badge.svg)
![python](https://github.com/miracl/oidc-samples/workflows/python/badge.svg)

This repository contains samples showcasing the integration between [MIRACL
Trust platform](https://miracl.com) and various OIDC libraries. There are also
[integration tests](./integration-tests/) that validate both the compatibility
and the correct operation of the client libraries.

## ENV Variables

All samples work with the following environment variables:

- `HOST` - Host to listen on. The default is "localhost".
- `PORT` - Port of the listening host. The default is "8000".
- `ISSUER` - OpenID Connect Issuer. The default is "https://api.mpin.io".
- `REDIRECT_URL` - The redirect URL of the app in the MIRACL Trust platform.
  The default value is "http://localhost:8000/login".
- `CLIENT_ID` - The Client ID of the app in the MIRACL Trust platform. It has
  no default value and is mandatory.
- `CLIENT_SECRET`- The Client Secret of the app in the MIRACL Trust platform.
  It has no default value and is mandatory.
- `PROXY_HOST`- The host address of the proxy behind which you run the sample.
  The default value is an empty string. It is used only when the setup requires
  a proxy, allowing us to validate that the OIDC client works behind a proxy.
- `PROXY_PORT`- The port of the proxy behind which you run the sample. The
  default value is an empty string. It is used only when the setup requires a
  proxy, allowing us to validate that the OIDC client works behind a proxy.

To get those values, you'll need to [register](https://miracl.com/resources/docs/get-started/register/)
and [create an app](https://miracl.com/resources/docs/get-started/low-code/) in
our platform.

## Usage

You can start every sample with its native tooling. For instructions, see the
README.md of the sample you are interested in.

Once you start your sample of choice, go to http://localhost:8000, which will
take you to the MIRACL Trust authorization page. You need to enrol the device
the first time you use the sample. Then, you can authenticate directly using M-PIN.

### Run with Docker

You can also use Docker to run any sample.

```bash
cd samples/<variant>
docker build -t sample .
docker run \
  --publish 8000:8000 \
  --env CLIENT_ID=<client-id> \
  --env CLIENT_SECRET=<client-secret> \
  sample
```

### Run on a different port with Docker

All samples use port 8000 as the default; hence, the sample starts at 8000, and
the default OIDC Redirect URL is http://localhost:8000/login. To change the
port you access the sample on, do the following:

```bash
docker run \
  --publish <custom-port>:8000 \
  --env CLIENT_ID=<client-id> \
  --env CLIENT_SECRET=<client-secret> \
  --env REDIRECT_URL=http://localhost:<custom_port>/login \
  sample
```

Note that we don't need to change the port the sample runs on in the container -
just the Docker mapping. The redirect URL must also be updated in
the command and the application configuration in the [MIRACL Trust
Portal](https://trust.miracl.cloud).

## Running through a proxy

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
