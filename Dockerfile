# Build traffic using go version 1.13.12
FROM golang:1.13.12 AS build

RUN apt-get update

WORKDIR /RBBN/proxy

COPY . .

RUN go build -o ./udp_proxy

##
FROM alpine:3.11.2

RUN set -x && \
    apk update && \
    apk add --no-cache bash \
              libc6-compat \
              openrc \
              curl \
              tcpdump \
              libstdc++ \
              iptables

WORKDIR /RBBN/proxy/

COPY --from=build /RBBN/proxy/udp_proxy  /RBBN/proxy
COPY --from=build /RBBN/proxy/assets /RBBN/proxy/assets
CMD ["/sbin/init"]


EXPOSE 8080 8000 61887/udp 61888/udp 500/udp 4500/udp
