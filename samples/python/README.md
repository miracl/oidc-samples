# MIRACL Trust Python OIDC Integration Sample

![python](https://github.com/miracl/oidc-samples/workflows/python/badge.svg)

This is an example of an
[OIDC integration](https://miracl.com/resources/docs/guides/authentication/oidc/)
with the [MIRACL Trust platform](https://miracl.com) using Python. It uses the
[IdPyOIDC](https://idpy-oidc.readthedocs.io/en/latest/) library to integrate
with the MIRACL Trust platform. Refer to the library's documentation for
detailed integration steps.

# Setup

To start an OIDC integration, you must create an OIDC application in the
[MIRACL Trust Portal](https://trust.miracl.com), as described in
[Start Low-Code Integration](https://miracl.com/resources/docs/get-started/low-code/).
The `Redirect URL` must be the same as the one the sample is run with. If you
use the sample's default value, it must be set to `http://localhost:8000/login`.

# Usage

To run the sample, you must first
[setup](https://packaging.python.org/en/latest/guides/installing-using-pip-and-virtual-environments/)
your environment.

```bash
cd samples/python
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt

python app.py \
    --client-id <client-id> \
    --client-secret <client-secret> \
    --issuer <oidc-issuer>
```

This starts the sample HTTP server. When you open http://localhost:8000/ in your
browser, you will be redirected to the MIRACL Trust authorization page to begin
the registration and authentication process. After successful authentication,
the sample retrieves the OIDC UserInfo endpoint and returns the result.

If you want to run the sample with Docker, follow the instructions in the main
[README](../../README.md#run-with-docker) file.
