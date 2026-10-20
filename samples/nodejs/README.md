# MIRACL Trust Node.js OIDC Integration Sample

![nodejs](https://github.com/miracl/oidc-samples/workflows/nodejs/badge.svg)

This is an example of an
[OIDC integration](https://miracl.com/resources/docs/guides/authentication/oidc/)
with the [MIRACL Trust platform](https://miracl.com) using Node.js. It uses the
[openid-client](https://www.npmjs.com/package/openid-client) library to
integrate with the MIRACL Trust platform. Refer to the library's documentation
for detailed integration steps.

# Setup

To start an OIDC integration:
1. Create a project in the [MIRACL Trust Portal](https://trust.miracl.com).
1. Go to the project’s OpenID Connect section to generate Client Credentials,
   as described in [Start Low-Code Integration](https://miracl.com/resources/docs/get-started/low-code/).

   You can find the OIDC Issuer in this section too.
1. In the Redirect URL field, add the sample’s Redirect URL. The default is:
   `http://localhost:8000/login`.
1. You must pass the app's credentials to the sample through environment
   variables as follows:

   ```bash
   export CLIENT_ID=<client-id>
   export CLIENT_SECRET=<client-secret>
   export ISSUER=<oidc-issuer>
   ```

# Usage

To run the sample, do the following:

```bash
cd samples/nodejs
npm install
node index.js
```

This starts the sample HTTP server. When you open http://localhost:8000/ in your
browser, you will be redirected to the MIRACL Trust authorization page to begin
the registration and authentication process. After successful authentication,
the sample retrieves the OIDC UserInfo endpoint and returns the result.

If you want to run the sample with Docker, follow the instructions in the main
[README](../../README.md#run-with-docker) file.
