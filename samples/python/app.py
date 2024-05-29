import os
import random

from flask import Flask, request
from oic.oic import Client, OicClientSettings
from oic.oic.message import AuthorizationResponse, RegistrationResponse
from oic.utils.authn.client import CLIENT_AUTHN_METHOD
from oic.utils.http_util import Redirect
from requests import Session
from werkzeug.middleware.proxy_fix import ProxyFix

state_storage = {}
session_storage = {}

proxy_host = os.environ.get("PROXY_HOST")
proxy_port = os.environ.get("PROXY_PORT")


def new_app():
    app = Flask(__name__)

    session = Session()
    oic_settings = OicClientSettings()

    if proxy_host is not None and proxy_port is not None:
        app.wsgi_app = ProxyFix(app.wsgi_app, x_for=1, x_proto=1, x_host=1, x_prefix=1)
        session.proxies = {"https": f"{proxy_host}:{proxy_port}"}
        oic_settings.requests_session=session

    client = Client(
        client_authn_method=CLIENT_AUTHN_METHOD,
        settings=(
            oic_settings if proxy_host is not None and proxy_port is not None else None
        ),
    )
    configuration_discovery = client.provider_config("https://api.mpin.io/")
    client.handle_provider_config(
        configuration_discovery, configuration_discovery["issuer"]
    )

    info = {
        "client_id": os.environ.get("CLIENT_ID"),
        "client_secret": os.environ.get("CLIENT_SECRET"),
    }
    client_reg = RegistrationResponse(**info)

    client.store_registration_info(client_reg)

    return app, client


app, client = new_app()


@app.route("/")
def index():
    state = str(random.randint(100000, 999999))
    nonce = str(random.randint(100000, 999999))
    state_storage[state] = {"state": state, "nonce": nonce}
    args = {
        "client_id": client.client_id,
        "response_type": "code",
        "scope": ["openid", "email", "profile"],
        "nonce": state_storage[state]["nonce"],
        "redirect_uri": "http://localhost:8000/login",
        "state": state_storage[state]["state"],
    }

    auth_req = client.construct_AuthorizationRequest(request_args=args)
    login_url = auth_req.request(client.authorization_endpoint)

    session_id = request.args.get("session")
    if session_id is None:
        return Redirect(login_url)

    user_info = session_storage[session_id]
    return user_info.to_json()


@app.route("/login")
def login():
    response = request.environ["QUERY_STRING"]
    aresp = client.parse_response(
        AuthorizationResponse, info=response, sformat="urlencoded"
    )
    code = aresp["code"]

    assert aresp["state"] in state_storage
    state_storage.pop(aresp["state"])

    args = {"code": code}

    client.do_access_token_request(
        state=aresp["state"], request_args=args, authn_method="client_secret_basic"
    )

    userinfo = client.do_user_info_request(method="GET", state=aresp["state"])

    session_id = str(random.randint(100000, 999999))
    session_storage[session_id] = userinfo

    return Redirect(f"/?session={session_id}")


if __name__ == "__main__":
    app.debug = True
    app.run(host="0.0.0.0", port=8000, debug=True)
