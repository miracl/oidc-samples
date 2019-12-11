# MIRACL Trust OIDC Integration Samples

![](https://github.com/miracl/oidc-samples/workflows/go/badge.svg)
![](https://github.com/miracl/oidc-samples/workflows/js/badge.svg)

## OIDC Credentials

To get client ID and secret check our documentation [here](https://docs.miracl.cloud/get-started/#client-id-and-secret).

## Run a sample as Docker container

```
cd samples/<variant>
docker build -t sample .
docker run -p 8000:8000 -e CLIENT_ID=<client-id> -e CLIENT_SECRET=<client-secret> sample
```

Afterwards visit `http://localhost:8000`. If you need to set another port, change the `8000:8000` port mapping and add an `ADDR` env variable after the client ID and secret in `docker run`.
