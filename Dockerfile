# Build
FROM golang:alpine AS builder
RUN apk --no-cache add git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/server ./cmd

# Runtime
FROM alpine:latest
RUN apk --no-cache add ca-certificates curl

# Install golang-migrate CLI
ARG MIGRATE_VERSION=v4.18.2
RUN curl -fsSL "https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz" \
  | tar xz -C /usr/local/bin

WORKDIR /app
COPY --from=builder /app/server .
COPY migrations ./migrations
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

EXPOSE 8081
ENTRYPOINT ["./entrypoint.sh"]
