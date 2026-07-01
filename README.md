# go-clean-rest-api

A small Go REST API starter with Chi, GORM, and PostgreSQL.

## Requirements

- Go 1.26+
- PostgreSQL

## Environment

Required:

- `DB_URL`: PostgreSQL connection string

Optional:

- `PORT`: HTTP port, defaults to `4545`

`.env` is loaded when present for local development, but it is not required in deployed environments.

## Run the server

```bash
go run ./cmd/server
```

## Run migrations

```bash
go run ./cmd/migrate
```

## Current routes

- `GET /health`

## Project layout

- `cmd/server`: API entrypoint
- `cmd/migrate`: schema migration entrypoint
- `config`: environment-backed app configuration
- `internal/auth`: auth-related persistence models
- `internal/server`: HTTP server and routing
- `internal/store`: database initialization and migration helpers
- `pkg/respond`: JSON response helpers
