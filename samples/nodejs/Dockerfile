FROM ghcr.io/miracl/oidc-samples/proxy:latest as proxy

FROM node:alpine
WORKDIR /src
COPY . .
ENV CA_CERT /src/ca.pem
COPY --from=proxy /src/ca.crt $CA_CERT
RUN npm install
EXPOSE 8000
ENTRYPOINT ["node", "/src/index.js"]
