using Microsoft.Owin;
using Microsoft.Owin.Security;
using Microsoft.Owin.Security.Cookies;
using Microsoft.Owin.Security.OpenIdConnect;
using Owin;
using System;
using System.Text;

[assembly: OwinStartup(typeof(dotnet.Startup))]
namespace dotnet
{
    public partial class Startup
    {
        public void Configuration(IAppBuilder app)
        {
            ConfigureAuth(app);
            Configure(app);
        }

        private void Configure(IAppBuilder app)
        {
            app.Map("/error", HandleError);
            app.Map("/logout", HandleLogout);

            app.Run(async context =>
            {
                if (context.Authentication.User == null || context.Authentication.User.Identity == null || !context.Authentication.User.Identity.IsAuthenticated)
                {
                    // This is what [Authorize] calls   
                    context.Authentication.Challenge(new AuthenticationProperties { RedirectUri = "/" }, OpenIdConnectAuthenticationDefaults.AuthenticationType);
                    return;
                }

                // Return json structured user info for our tests.
                context.Response.ContentType = "text/json";
                StringBuilder resp = new StringBuilder("{");
                foreach (var claim in context.Authentication.User.Claims)
                {
                    resp.Append(string.Format("\"{0}\":\"{1}\",", claim.Type, claim.Value));
                }

                resp.Remove(resp.Length - 1, 1);
                resp.Append("}");
                await context.Response.WriteAsync(resp.ToString());
            });
        }

        private void HandleError(IAppBuilder app)
        {
            app.Run(async context =>
            {
                string msg = context.Request.Query["message"];
                await context.Response.WriteAsync("Error occurred: " + msg);
            });
        }

        private void HandleLogout(IAppBuilder app)
        {
            app.Run(async context =>
            {
                if (context.Authentication.User != null && context.Authentication.User.Identity != null && context.Authentication.User.Identity.IsAuthenticated)
                {
                    context.Authentication.SignOut(OpenIdConnectAuthenticationDefaults.AuthenticationType, CookieAuthenticationDefaults.AuthenticationType);
                }
               
                await context.Response.WriteAsync("You've beed just logout!");
            });
        }
    }
}