"""
This example integrates the IdPyOIDC library with the MIRACL Trust platform.
"""

import argparse
import os
import random
from urllib.parse import parse_qs

from flask import Flask, redirect, request
from idpyoidc.client.rp_handler import RPHandler

session_storage = {}
proxy_host = os.environ.get("PROXY_HOST")
proxy_port = os.environ.get("PROXY_PORT")

parser = argparse.ArgumentParser()
parser.add_argument("-cid", "--client-id", dest="client_id", help="Client ID")
parser.add_argument(
    "-cs", "--client-secret", dest="client_secret", help="Client Secret"
)
parser.add_argument("-iss", "--issuer", dest="issuer", help="OIDC Issuer")

args = parser.parse_args()


def new_app():
    """
    Creates an instance of a flask application, relying party handler and a relying party client.
    """
    application = Flask(__name__)

    if args.client_id:
        client_id = args.client_id
    else:
        client_id = os.environ.get("CLIENT_ID")

    if args.client_secret:
        client_secret = args.client_secret
    else:
        client_secret = os.environ.get("CLIENT_SECRET")

    if args.issuer:
        issuer = args.issuer
    elif os.environ.get("ISSUER"):
        issuer = os.environ.get("ISSUER")
    else:
        issuer = "https://api.mpin.io/"

    client_configs = {
        "miracl": {
            "issuer": issuer,
            "client_id": client_id,
            "client_secret": client_secret,
            "client_type": "oidc",
            "redirect_uris": ["http://localhost:8000/login"],
        }
    }
    rp_handler = RPHandler(
        client_configs=client_configs,
        httpc_params=(
            {
                "proxies": {
                    "http": f"{proxy_host}:{proxy_port}",
                    "https": f"{proxy_host}:{proxy_port}",
                }
            }
            if proxy_host is not None and proxy_port is not None
            else None
        ),
    )
    rp_client = rp_handler.init_client("miracl")
    rp_handler.do_provider_info(client=rp_client)

    return application, rp_handler, rp_client


app, rph, client = new_app()


@app.route("/")
def index():
    """
    Makes the request to the authorization endpoint, returns the user info if session id is found.
    """
    res = rph.init_authorization(client=client, req_args={"scope": ["openid", "email"]})
    session_id = request.args.get("session")
    if session_id is None:
        return redirect(res)

    user_info = session_storage[session_id]

    return user_info.to_json()


@app.route("/login")
def login():
    """
    Makes the access token and user info requests.
    """
    response = parse_qs(request.environ["QUERY_STRING"])
    authorization_code = response["code"][0]
    authorization_state = response["state"][0]
    response_dict = {"code": authorization_code, "state": authorization_state}
    final_auth = rph.finalize_auth(
        client=client, issuer="miracl", response=response_dict
    )
    token = rph.get_access_and_id_token(
        authorization_response=final_auth, client=client
    )
    user_info = rph.get_user_info(
        state=authorization_state, client=client, access_token=token["access_token"]
    )
    session_id = str(random.randint(100000, 999999))
    session_storage[session_id] = user_info

    return redirect(f"/?session={session_id}")


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)
