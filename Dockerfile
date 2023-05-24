## Build
FROM golang:1.20-buster AS build


COPY . /usr/src/bup

WORKDIR /usr/src/bup

RUN go build -o /usr/local/bin/bup


## Deploy
FROM debian:stable-slim

RUN useradd bup
COPY --from=build /usr/local/bin/bup /usr/local/bin/bup

WORKDIR /opt/bup

RUN chown bup -R /opt/bup

USER bup

ENTRYPOINT ["/usr/local/bin/bup"]
