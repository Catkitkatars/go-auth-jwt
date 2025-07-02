FROM golang:1.24 AS build
WORKDIR /src
COPY .. .
RUN --mount=type=ssh --mount=type=cache,target="/go/pkg/mod" \
    go mod download
RUN --mount=type=cache,target="/root/.cache/go-build" \
    --mount=type=cache,target="/go/pkg/mod" \
    CGO_ENABLED=1 go build -o auth-jwt cmd/app/main.go

FROM alpine:3.20

WORKDIR /auth-jwt

RUN apk add --no-cache \
    mysql-client \
    postgresql-client \
    ca-certificates \
    libc6-compat \
    && mkdir -p /auth-jwt/logs \
    && mkdir -p /auth-jwt/migrations

COPY --from=build /src/auth-jwt /auth-jwt
COPY --from=build /src/.env /auth-jwt/.env

COPY internal/store/migrations /auth-jwt/migrations

EXPOSE 8080
ENTRYPOINT ["./auth-jwt"]