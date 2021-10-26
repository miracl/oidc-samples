# syntax = docker/dockerfile:1.0-experimental
FROM ubuntu:latest AS c_builder

ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    cmake \
    g++ \
    gcc \
    git \
    make

ARG REVISION=master
ARG CRYPTO_GIT_USER

RUN --mount=type=secret,id=crypto_git_token,required \
    echo Building branch ${REVISION} \
    && git clone https://${CRYPTO_GIT_USER}:$(cat /run/secrets/crypto_git_token)@gitlab.corp.miracl.com/mfa/pkg/crypto.git \
    && cd crypto \
    && git checkout ${REVISION} \
    && make default \
    AMCL_TEST=OFF \
    AMCL_CURVE=ANSSI,BLS383,BLS461,BN254,BN254CX,BRAINPOOL,C41417,ED25519,FP256BN,FP512BN,GOLDILOCKS,HIFIVE,NIST256,NIST384,NIST521,NUMS256E,NUMS256W,NUMS384E,NUMS384W,NUMS512E,NUMS512W \
    AMCL_RSA=2048,3072,4096 \
    AMCL_BUILD_DOXYGEN=OFF \
    && cd target/default \
    && make \
    && make install

FROM registry.corp.miracl.com/gitlab-ci/snippets/golang-ci:latest

COPY --from=c_builder /opt/amcl /opt/amcl

ENV LD_LIBRARY_PATH="/opt/amcl/lib"
ENV CGO_LDFLAGS="-L /opt/amcl/lib"
ENV CGO_CPPFLAGS="-I /opt/amcl/include"
