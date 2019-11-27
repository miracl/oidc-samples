# OIDC Go Sample

![](https://github.com/ivivanov/oidc-samples/workflows/Go/badge.svg)

## Demo

The demo application is in `cmd/demo`

### Options

- `client-id` - The client id, registered in Miracl OIDC provider
- `client-secret`- The corresponding client secret
- `redirect-url` - The registered redirect URL
- `addr` - Host to bind and port to listen on in the form host:port; the default is ":8000" which means bind all available interfaces and listen on port 8000
- `templates-dir` - Folder holding the templates - absolute or relative to binary
- `issuer` - Backend url to the token issuer

### Build and run as Docker container

```
docker build -t oidc-go-sample oidc-go-sample/cmd/demo/Dockerfile
docker run -p 8000:8000 oidc-go-sample:latest -client-id <client-id> -client-secret <client-secret>
```
