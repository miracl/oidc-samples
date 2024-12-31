# MIRACL Trust DotNet Framework 4.8 OIDC Integration Sample

![dotnet](https://github.com/miracl/oidc-samples/workflows/dotnet/badge.svg)


This is an example of an [OIDC integration](https://miracl.com/resources/docs/guides/authentication/oidc/)
with the [MIRACL Trust platform](https://miracl.com) using standard dotnet framework library.

# Dependencies

This sample uses Owin and OpenID Connect dependencies to integrate with the MIRACL Trust platform:
- [https://www.nuget.org/packages/Microsoft.AspNet.Identity.Owin](https://www.nuget.org/packages/Microsoft.AspNet.Identity.Owin/)
- [https://www.nuget.org/packages/Microsoft.Owin.Host.SystemWeb](https://www.nuget.org/packages/Microsoft.Owin.Host.SystemWeb)
- [https://www.nuget.org/packages/OpenAthens.Owin.Security.OpenIdConnect](https://www.nuget.org/packages/OpenAthens.Owin.Security.OpenIdConnect/)

# Setup

The sample can be run directly with dotnet or with a Docker container. If you choose
the latter, follow the instructions [here](../../README.md).

To start an OIDC integration, you must create an OIDC application in the
[MIRACL Trust Portal](https://trust.miracl.com) as described [here](https://miracl.com/resources/docs/get-started/low-code/).
The `Redirect URL` must be the same as the one the sample is run with. If you use the
sample's default value, it must be set to `http://localhost:59504/login`.
You must pass the app's credentials to the sample through its `web.config` file
as follows:

``` bash
  <appSettings>
    .....
    <add key="REDIRECT_URL" value="http://localhost:59504/login" />
    <add key="CLIENT_ID" value="CLIENT_ID" />
    <add key="CLIENT_SECRET" value="CLIENT_SECRET" />
  </appSettings>
```

# Usage

To run the sample, do the following:

* open the OidcSample.sln in your Visual Studio
* right click over the solution in the Solution Explorer and press
  `Restore NuGet Packages`
* run the sample using `F5` or the run button of the redactor.

This runs a sample HTTP server which navigates you to the MIRACL Trust authorization
page to start the registration and authentication.
