# MIRACL Trust Python OIDC Integration Sample

![python](https://github.com/miracl/oidc-samples/workflows/python/badge.svg)

This is an example of an [OIDC integration](https://miracl.com/resources/docs/guides/authentication/oidc/)
with the [MIRACL Trust platform](https://miracl.com) using Python.

# Dependencies

This sample uses [IdPyOIDC](https://idpy-oidc.readthedocs.io/en/latest/) library
to integrate with the MIRACL Trust platform.

# Setup

The sample can be run directly with Python or with a Docker container. If you choose
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

To run the sample, you need first to [setup](https://packaging.python.org/en/latest/guides/installing-using-pip-and-virtual-environments/)
your environment.

```bash
cd samples/python
python3 -m venv .venv
source .venv/bin/activate
python3 -m pip install -r requirements.txt

python app.py --client-id CLIENTID --client-secret CLIENTSECRET
```

This runs a sample HTTP server and when you open http://localhost:8000/,
it navigates you to the MIRACL Trust authorization page to start the
registration and authentication.
