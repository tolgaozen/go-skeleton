FROM golang:1.24.3-alpine3.20@sha256:9f98e9893fbc798c710f3432baa1e0ac6127799127c3101d2c263c3a954f0abe as skeleton-builder
WORKDIR /go/src/app
RUN apk update && apk add --no-cache git
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,target=/go/pkg/mod CGO_ENABLED=0 go build -v ./cmd/skeleton/

FROM cgr.dev/chainguard/static:latest@sha256:b2e1c3d3627093e54f6805823e73edd17ab93d6c7202e672988080c863e0412b
COPY --from=ghcr.io/grpc-ecosystem/grpc-health-probe:v0.4.28 /ko-app/grpc-health-probe /usr/local/bin/grpc_health_probe
COPY --from=skeleton-builder /go/src/app/skeleton /usr/local/bin/skeleton
ENV PATH="$PATH:/usr/local/bin"
ENTRYPOINT ["skeleton"]
CMD ["serve"]
