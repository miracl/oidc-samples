using IdentityModel.Client;
using Microsoft.IdentityModel.Protocols.OpenIdConnect;
using Microsoft.IdentityModel.Tokens;
using Microsoft.Owin.Security;
using Microsoft.Owin.Security.Cookies;
using Microsoft.Owin.Security.Notifications;
using Microsoft.Owin.Security.OpenIdConnect;
using Owin;
using System;
using System.Configuration;
using System.Net.Http;
using System.Threading.Tasks;

namespace dotnet
{
    public partial class Startup
    {
        public static string OidcAuthority = Environment.GetEnvironmentVariable("ISSUER") ?? ConfigurationManager.AppSettings["ISSUER"];
        public static string OidcRedirectUrl = Environment.GetEnvironmentVariable("REDIRECT_URL") ?? ConfigurationManager.AppSettings["REDIRECT_URL"];
        public static string OidcClientId = Environment.GetEnvironmentVariable("CLIENT_ID") ?? ConfigurationManager.AppSettings["CLIENT_iD"];
        public static string OidcClientSecret = Environment.GetEnvironmentVariable("CLIENT_SECRET") ?? ConfigurationManager.AppSettings["CLIENT_SECRET"];
        public const string NameIdentifier = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/nameidentifier";
        
        public void ConfigureAuth(IAppBuilder app)
        {
            app.SetDefaultSignInAsAuthenticationType(CookieAuthenticationDefaults.AuthenticationType);
            app.UseCookieAuthentication(new CookieAuthenticationOptions());

            var oidcOptions = new OpenIdConnectAuthenticationOptions
            {
                Authority = OidcAuthority,
                ClientId = OidcClientId,
                ClientSecret = OidcClientSecret,
                ResponseType = OpenIdConnectResponseType.Code,
                SaveTokens = true,
                RedeemCode = true,
                TokenValidationParameters = new TokenValidationParameters
                {
                    NameClaimType = NameIdentifier
                },
                PostLogoutRedirectUri = OidcRedirectUrl,
                RedirectUri = OidcRedirectUrl,
                Scope = "openid email profile",
                UsePkce = false,
                Notifications = new OpenIdConnectAuthenticationNotifications()
                {
                    AuthorizationCodeReceived = OnAuthorizationCodeReceived,
                    AuthenticationFailed = OnAuthenticationFailed,
                    SecurityTokenValidated = OnSecurityTokenValidated
                }
            };

            app.UseOpenIdConnectAuthentication(oidcOptions);
        }

        private Task OnAuthenticationFailed(AuthenticationFailedNotification<OpenIdConnectMessage, OpenIdConnectAuthenticationOptions> context)
        {
            context.HandleResponse();
            context.Response.Redirect("/error?message=" + context.Exception.Message);
            return Task.FromResult(0);
        }

        private async Task OnAuthorizationCodeReceived(AuthorizationCodeReceivedNotification c)
        {
            // If there is a code in the OpenID Connect response, redeem it for an access token and refresh token, and store those away.
            HttpClient client = new HttpClient();
            var response = await client.RequestAuthorizationCodeTokenAsync(new AuthorizationCodeTokenRequest
            {
                Address = "https://api.mpin.io/oidc/token",

                ClientId = c.TokenEndpointRequest.ClientId,
                ClientSecret = c.TokenEndpointRequest.ClientSecret,

                Code = c.TokenEndpointRequest.Code,
                RedirectUri = c.TokenEndpointRequest.RedirectUri,
            });

            if (response.IsError)
            {
                c.Response.Redirect("/error?message=" + response.Error);
            }

            c.HandleCodeRedemption(response.AccessToken, response.IdentityToken);
        }

        private Task OnSecurityTokenValidated(SecurityTokenValidatedNotification<OpenIdConnectMessage, OpenIdConnectAuthenticationOptions> context)
        {
            var identity = context.AuthenticationTicket.Identity;
            string nameid = identity.FindFirst(NameIdentifier).Value;
            identity.AddClaim(new System.Security.Claims.Claim("email", nameid));
            identity.AddClaim(new System.Security.Claims.Claim("sub", nameid));
            return Task.CompletedTask;
        }
    }
}