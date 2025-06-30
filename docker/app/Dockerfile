FROM golang:1.24 AS build
WORKDIR /src
COPY .. .
RUN --mount=type=ssh --mount=type=cache,target="/go/pkg/mod" \
    go mod download
RUN --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=cache,target="/go/pkg/mod" \
    CGO_ENABLED=1 go build -o auth-jwt cmd/app/main.go

FROM ubuntu:22.04
WORKDIR /auth-jwt

RUN apt-get update && apt-get install -y \
    default-mysql-client \
    postgresql-client && \
    rm -rf /var/lib/apt/lists/* && \
    mkdir -p /auth-jwt/logs \
    && mkdir -p /auth-jwt/migrations

COPY --from=build /src/auth-jwt /auth-jwt
COPY --from=build /src/.env /auth-jwt/.env

COPY internal/store/migrations /auth-jwt/migrations

EXPOSE 8080
ENTRYPOINT ["./auth-jwt"]