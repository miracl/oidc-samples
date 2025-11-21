# MIRACL Proxy

Proxy implementation that allows modification of the traffic from and to the
samples for testing purposes.

## Modes

The proxy has two modes - one that forwards the traffic and one that allows us
to modify the requests and responses coming and going through the proxy, using
Man-In-The-Middle (MITM).

### MITM

A POST request to the `/session` endpoint starts an MITM session. The body of
the request must contain a JSON document with `modifyUrl` property. This
property is the URL where the proxy will redirect all traffic until the session
is stopped. A `DELETE` request to the `/session` endpoint stops the session.

## Usage

The proxy is written in Go, so it can be run either as a Go service:

```bash
cd ./proxy
go run .
```

or as a Docker container:

```
docker build -t proxy .
docker run --publish 8080:8080 proxy
```

Then, you must pass the PROXY_HOST and PROXY_PORT environment variables to your
sample as described in the
[Run through a proxy](../README.md?tab=readme-ov-file#run-through-proxy) section
of the main README.
