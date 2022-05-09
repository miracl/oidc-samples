# MIRACL Proxy

Proxy implementation that allows modification of the traffic from and to the
samples for testing purposes.

## Modes

The proxy has two modes - one that simply forwards the traffic and one which
allows us to modify the requests and responses coming and going through the
proxy, using Man-In-The-Middle (MITM).

### MITM

A POST request to the `/session` endpoint starts a MITM session. The body of the
request should contain JSON document with `modifyUrl` property. This property
is the URL where the proxy will redirect all traffic until the session is
stopped. A `DELETE` request to the `/session` endpoint stops the session.
