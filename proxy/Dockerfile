FROM golang:latest
WORKDIR /src
COPY . .
RUN go build -o /usr/local/bin/proxy .
ENTRYPOINT ["proxy"]
