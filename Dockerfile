FROM golang:alpine AS build

WORKDIR /tmp/build
COPY go.mod go.sum ./
RUN go mod download -x

ARG CGO_ENABLED=0
ADD . .
RUN go build -v -trimpath

# ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

FROM spritsail/alpine:3.20

LABEL org.opencontainers.image.authors="Joe Groocock <jool-exporter@frebib.net>" \
      org.opencontainers.image.title="Prometheus Jool NAT64 exporter" \
      org.opencontainers.image.url="https://github.com/frebib/jool-exporter" \
      org.opencontainers.image.description="Prometheus exporter for https://github.com/NICMx/Jool"

RUN apk add --no-cache jool-tools
COPY --from=build /tmp/build/jool-exporter /jool-exporter

EXPOSE 9441
ENTRYPOINT ["/jool-exporter"]
