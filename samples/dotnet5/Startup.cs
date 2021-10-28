using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.IdentityModel.Protocols.OpenIdConnect;
using Microsoft.AspNetCore.Authentication.OpenIdConnect;
using Microsoft.AspNetCore.Authentication.Cookies;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Http;
using System.Net.Http;
using Microsoft.Extensions.Options;
using Microsoft.AspNetCore.Authentication;
using System.Text.Encodings.Web;
using IdentityModel.Client;
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
                options.Scope.Add("openid email profile");

                options.UsePkce = false;
                options.SaveTokens = true;

                options.Events = new OpenIdConnectEvents()
                {
                    OnAuthorizationCodeReceived = async c =>
                    {
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

            app.UseCookiePolicy();
            app.UseAuthentication();
            app.UseAuthorization();

            app.Run(async context =>
            {
                var response = context.Response;

                // This is what [Authorize] calls
                var userResult = await context.AuthenticateAsync();
                var user = userResult.Principal;
                var props = userResult.Properties;

                // This is what [Authorize(ActiveAuthenticationSchemes = OpenIdConnectDefaults.AuthenticationScheme)] calls
                // var user = await context.AuthenticateAsync(OpenIdConnectDefaults.AuthenticationScheme);

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

        private static string HtmlEncode(string content) =>
            string.IsNullOrEmpty(content) ? string.Empty : HtmlEncoder.Default.Encode(content);

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
