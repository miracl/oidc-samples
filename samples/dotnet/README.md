# MIRACL Trust .NET Framework 4.8 OIDC Integration Sample

![dotnet](https://github.com/miracl/oidc-samples/workflows/dotnet/badge.svg)

This is an example of an
[OIDC integration](https://miracl.com/resources/docs/guides/authentication/oidc/)
with the [MIRACL Trust platform](https://miracl.com) using standard .NET
framework library. It uses the Owin and OpenID Connect dependencies to integrate
with the MIRACL Trust platform:

- [https://www.nuget.org/packages/Microsoft.AspNet.Identity.Owin](https://www.nuget.org/packages/Microsoft.AspNet.Identity.Owin/)
- [https://www.nuget.org/packages/Microsoft.Owin.Host.SystemWeb](https://www.nuget.org/packages/Microsoft.Owin.Host.SystemWeb)
- [https://www.nuget.org/packages/OpenAthens.Owin.Security.OpenIdConnect](https://www.nuget.org/packages/OpenAthens.Owin.Security.OpenIdConnect/)

# Setup

To start an OIDC integration, you must create an OIDC application in the
[MIRACL Trust Portal](https://trust.miracl.com), as described in
[Start Low-Code Integration](https://miracl.com/resources/docs/get-started/low-code/).
The `Redirect URL` must be the same as the one the sample is run with. If you
use the sample's default value, it must be set to
`http://localhost:59504/login`. You must pass the app's credentials to the
sample through its `web.config` file as follows:

```bash
  <appSettings>
    .....
    <add key="REDIRECT_URL" value="http://localhost:59504/login" />
    <add key="CLIENT_ID" value="CLIENT_ID" />
    <add key="CLIENT_SECRET" value="CLIENT_SECRET" />
    <add key="ISSUER" value="OIDC_ISSUER" />
  </appSettings>
```

# Usage

To run the sample, follow these steps:

1. Open the OidcSample.sln in your Visual Studio.
1. Right-click over the solution in the Solution Explorer and select **Restore
   NuGet Packages**.
1. Press **F5** or click **Run** in the editor to run the sample.

This starts the sample HTTP server. When you open http://localhost:8000/ in your
browser, you will be redirected to the MIRACL Trust authorization page to begin
the registration and authentication process. After successful authentication,
the sample retrieves the OIDC UserInfo endpoint and returns the result.

If you want to run the sample with Docker, follow the instructions in the main
[README](../../README.md#run-with-docker) file.
