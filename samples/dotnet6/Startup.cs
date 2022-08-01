using IdentityModel.Client;
using Microsoft.AspNetCore.Authentication;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Authentication.OpenIdConnect;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Text;

namespace OidcSample
{
    public class Startup
    {
        private const string MiraclIssuer = "https://api.mpin.io";

        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            HttpClientHandler clientHandler = null;
            if (!string.IsNullOrEmpty(this.Configuration["PROXY_HOST"]))
            {
                int port;
                if (!int.TryParse(this.Configuration["PROXY_PORT"] ?? "8080", out port))
                {
                    throw new ArgumentException($"HOST_PORT should be a valid port number in the integer range. The current value {this.Configuration["PROXY_PORT"]} cannot be parsed.");
                }
                var proxy = new WebProxy(this.Configuration["PROXY_HOST"], port);
                proxy.UseDefaultCredentials = true;
                proxy.BypassProxyOnLocal = true;

                clientHandler = new HttpClientHandler
                {
                    Proxy = proxy,
                    UseProxy = true,
                    ServerCertificateCustomValidationCallback = (o, c, ch, er) => true
                };
            }

            services.AddRazorPages();

            services.AddAuthentication(options =>
            {
                options.DefaultAuthenticateScheme = OpenIdConnectDefaults.AuthenticationScheme;
                options.DefaultChallengeScheme = OpenIdConnectDefaults.AuthenticationScheme;
                options.DefaultScheme = CookieAuthenticationDefaults.AuthenticationScheme;
            })
            .AddCookie()
            .AddOpenIdConnect(OpenIdConnectDefaults.AuthenticationScheme, "MIRACL", options =>
            {
                if (clientHandler != null)
                {
                    options.BackchannelHttpHandler = clientHandler;
                }

                options.NonceCookie.SameSite = SameSiteMode.Lax;
                options.CorrelationCookie.SameSite = SameSiteMode.Lax;

                options.Authority = this.Configuration["ISSUER"] ?? MiraclIssuer;
                options.ClientId = this.Configuration["CLIENT_ID"];
                options.ClientSecret = this.Configuration["CLIENT_SECRET"];

                options.ResponseType = "code";
                options.SaveTokens = true;
                options.GetClaimsFromUserInfoEndpoint = true;
                options.AuthenticationMethod = OpenIdConnectRedirectBehavior.RedirectGet;
                options.CallbackPath = this.GetCallbackPath();

                options.Scope.Clear();
                options.Scope.Add("openid email");

                options.UsePkce = false;
                options.SaveTokens = true;
                options.ClaimActions.MapJsonKey("email", "email", "url");
                options.ClaimActions.MapJsonKey("sub", "sub", "string");

                options.Events = new OpenIdConnectEvents()
                {
                    OnAuthorizationCodeReceived = async c =>
                    {
                        var client = clientHandler != null ? new HttpClient(clientHandler) : new HttpClient();
                        var response = await client.RequestAuthorizationCodeTokenAsync(new AuthorizationCodeTokenRequest
                        {
                            Address = MiraclIssuer + "/oidc/token",

                            ClientId = c.TokenEndpointRequest.ClientId,
                            ClientSecret = c.TokenEndpointRequest.ClientSecret,

                            Code = c.TokenEndpointRequest.Code,
                            RedirectUri = c.TokenEndpointRequest.RedirectUri,
                        });

                        if (response.IsError)
                        {
                            throw new Exception(response.Error);
                        }

                        c.HandleCodeRedemption(response.AccessToken, response.IdentityToken);
                    }
                };
            });
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            app.UseDeveloperExceptionPage();

            app.UseAuthentication();
            app.UseAuthorization();

            app.Map("/logout", HandleLogout);

            app.Run(async context =>
            {
                var response = context.Response;

                // This is what [Authorize] calls
                var userResult = await context.AuthenticateAsync();
                var user = userResult.Principal;
                var props = userResult.Properties;

                // Not authenticated
                if (user == null || !user.Identities.Any(identity => identity.IsAuthenticated))
                {
                    // This is what [Authorize] calls
                    await context.ChallengeAsync();

                    return;
                }


                // Return json structured user info for our tests.
                response.ContentType = "text/json";
                StringBuilder rr = new StringBuilder("{");
                foreach (var claim in context.User.Claims)
                {
                    rr.Append(string.Format("\"{0}\":\"{1}\",", claim.Type, claim.Value));
                }

                rr.Remove(rr.Length - 1, 1);
                rr.Append("}");
                await response.WriteAsync(rr.ToString());
            });
        }

        private void HandleLogout(IApplicationBuilder app)
        {
            app.Run(async context =>
            {
                if (context.User != null && context.User.Identity != null && context.User.Identity.IsAuthenticated)
                {
                    await context.SignOutAsync();
                }

                await context.Response.WriteAsync("You've been just logout!");
            });
        }

        private string GetCallbackPath()
        {
            if (string.IsNullOrEmpty(this.Configuration["REDIRECT_URL"]))
            {
                return "/login";
            }

            string redirectUrl = this.Configuration["REDIRECT_URL"];
            string callbackPath = redirectUrl.Remove(0, redirectUrl.IndexOf("://") + 3);
            callbackPath = callbackPath.Remove(0, callbackPath.IndexOf("/"));

            return callbackPath;
        }
    }
}
