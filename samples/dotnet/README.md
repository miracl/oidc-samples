# MIRACL Trust DotNet Framework 4.8 OIDC Integration Sample

This is an example integration with the standard dotnet framework library using Owin and OpenID Connect dependencies: 
- [https://www.nuget.org/packages/Microsoft.AspNet.Identity.Owin/](https://www.nuget.org/packages/Microsoft.AspNet.Identity.Owin/)
- [https://www.nuget.org/packages/Microsoft.Owin.Host.SystemWeb](https://www.nuget.org/packages/Microsoft.Owin.Host.SystemWeb)
- [https://www.nuget.org/packages/OpenAthens.Owin.Security.OpenIdConnect/](https://www.nuget.org/packages/OpenAthens.Owin.Security.OpenIdConnect/)

## Usage

You can run any sample as Docker container

```
cd samples/<variant>
docker build -t sample .
docker run -p 8000:8000 -e CLIENT_ID=<client-id> -e CLIENT_SECRET=<client-secret> sample
```
