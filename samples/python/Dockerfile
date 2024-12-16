ARG PYTHON_VERSION=3.12.3
FROM ghcr.io/miracl/oidc-samples/proxy:latest as proxy

FROM python:${PYTHON_VERSION}-slim as base

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
WORKDIR /src
COPY . .

RUN --mount=type=cache,target=/root/.cache/pip \
    --mount=type=bind,source=requirements.txt,target=requirements.txt \
    python -m pip install -r requirements.txt

COPY --from=proxy /src/ca.crt /usr/local/share/ca-certificates/proxy.crt
RUN update-ca-certificates

EXPOSE 8000
CMD python app.py
