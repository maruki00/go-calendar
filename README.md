# Go Calendar

A lightweight Go-based calendar backend + admin UI using Gin and SQLite.

## Project Overview

This repository implements an event calendar service with endpoints for creating, updating, deleting, and getting events, including common event operations. It serves a static frontend dashboard from `cmd/index.html` and assets under `db/static`.

### Key features
- REST-like routes for event operations
- SQLite persistence via `pkg/postgres` (and alternate handlers in `pkg/sqlite`)
- Clean domain structure under `internal/event` (controllers, services, repositories)
- Admin UI with FullCalendar and AdminLTE assets

## Architecture

- `cmd/main.go` — Application entrypoint and HTTP route definitions
- `internal/event/userinterface/controllers` — HTTP handlers
- `internal/event/app/services` — Business logic service layer
- `internal/event/infra/Repositories` — Database repository implementations
- `pkg/postgres` — DB handler and connection logic
- `db/static` — Static assets for calendar UI

## Quick Start

### Requirements
- Go 1.23+
- Linux/macOS/Windows

### Run locally

```bash
cd /home/user/dev/go-calendar
go mod tidy
go build -o calendar ./cmd
go run ./calendar
```

Open http://localhost:5600 in your browser.

## API Endpoints

The app runs on port `5600` by default.

- `POST /api/v1/event/create` - Create an event
- `POST /api/v1/event/delete` - Delete an event
- `POST /api/v1/event/update` - Update an event
- `POST /api/v1/events` - Get events
- `POST /api/v1/event/common/create` - Create common event
- `POST /api/v1/event/common/delete` - Delete common event
- `GET /` - Serve main UI (`index.html`)

### Example JSON payload (Create)
```json
{
  "name": "Team Meeting",
  "description": "Weekly sync",
  "start_date": "2026-03-15T10:00:00Z",
  "end_date": "2026-03-15T11:00:00Z",
  "user_id": "1"
}
```

## Database

The server initializes `./db.db` via `pkg/postgres.NewDBHandler("./db.db")`. If missing, the app creates it.

## Frontend

The root path renders `cmd/index.html` using Go templates and static files are served at `/static` from `/opt/calendar/static` (via `server.Static("/static", "/opt/calendar/static")`).

> Note: For local development, ensure this path exists or update the static path in `cmd/main.go`.

## Project Commands

- `go run ./cmd` — Start server
- `go test ./...` — Run all tests (if any are added)

## Contribution

1. Fork the repository.
2. Create a feature branch.
3. Add tests and documentation.
4. Open a pull request.

