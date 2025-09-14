# Builder stage
FROM golang:1.25-alpine AS builder

ARG GIT_ACCESS_TOKEN

RUN --mount=type=cache,target=/var/cache/apk \
  echo "edge" > /etc/alpine-release && \
  apk update && \
  apk upgrade && \
  apk add build-base make libcap musl-dev coreutils tar bash git xxd xz sqlite zig

# Enable CGO and set platform
ENV CGO_ENABLED=1 GOOS=linux CC="zig cc"

RUN git config --global url."https://${GIT_ACCESS_TOKEN}:x-oauth-basic@repo.reapertech.com/".insteadOf "https://repo.reapertech.com/"

WORKDIR /app

# Copy go mod and sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with CGO
RUN go build -o zeno -ldflags "-linkmode=external -extldflags=-static" .

FROM scratch AS artifact
ARG TARGETARCH

COPY --from=builder /app/zeno /zeno-static-$TARGETARCH

# Final smaller runtime stage
FROM alpine:3.22 AS release

# Install runtime dependencies: C++ standard libs and SSL certs
RUN apk add --no-cache libstdc++ libgcc ca-certificates bash

WORKDIR /app

COPY ./entrypoint.sh /entrypoint.sh
# Copy built binary from builder
COPY --from=builder /app/zeno .

# Expose port
EXPOSE 8080

# Start app
CMD ["/entrypoint.sh"]
