# Social Network API (WIP)

This is a backend API built using Go, designed to serve as the foundation of a social networking application. The current implementation includes basic configuration, logging, error handling, and a health check endpoint.

---

## Features Implemented So Far

* Configuration through command-line flags
* Structured logging using Go's `log/slog` package
* Centralized error handling with contextual logging
* JSON-based error and success responses
* Health check endpoint at `/healthcheck`

---

## Project Structure

```
social-network/
├── cmd/
│   └── api/
│       ├── main.go            # Entry point: config, logger, HTTP server setup
│       ├── handlers.go        # HTTP handlers (e.g., /healthcheck)
│       ├── helpers.go         # Error response helpers, log helpers
│       └── response.go        # JSON response writer
├── go.mod
├── go.sum
```

---

## How to Run the Application

Make sure you have Go installed (version 1.21 or later recommended).

```bash
# Run the application
$ go run ./cmd/api/ -port=4000 -env=development
```

> The server will start on the specified port (default 4000). You can visit `http://localhost:4000/healthcheck` in your browser or use curl to test the endpoint:
>
> ```bash
> curl http://localhost:4000/healthcheck
> ```

---

## Configuration Flags

| Flag    | Description                              | Default       |
| ------- | ---------------------------------------- | ------------- |
| `-port` | Port number the server listens on        | `4000`        |
| `-env`  | Application environment (dev/stage/prod) | `development` |

---

## Endpoints

### GET /healthcheck

**Purpose:** Check the availability and environment of the running server.

**Sample Response:**

```json
{
  "status": "available",
  "environment": "development"
}
```

---

## Error Handling

The application uses centralized, consistent JSON error responses. If an error occurs (e.g., unexpected server issue, 404 not found, or unsupported method), the response will include an `error` field with a descriptive message.

### Example: 500 Internal Server Error

```json
{
  "error": "The server encountered a problem and could not process your request"
}
```

### Example: 404 Not Found

```json
{
  "error": "The requested resource could not be found"
}
```

### Example: 405 Method Not Allowed

```json
{
  "error": "the POST method is not supported for this resource"
}
```

All errors are also logged with context (HTTP method and URI) using the logger:

```
level=ERROR msg="sql: no rows in result set" method=GET uri=/users/42
```

---

## Logging

Logging is handled using Go's standard `log/slog` package. The logger outputs structured key-value pairs and includes the following context by default:

* Timestamp
* Log level (INFO, ERROR)
* Message
* Method and URI for HTTP errors

Example log entry:

```
time=2025-07-12T13:00:00.123+03:00 level=INFO msg="starting server" addr=":4000" env="development"
```

---

## Next Steps

* Add more routes for authentication and other functionalities i.e posts
* Implement middleware (e.g., request logging)
* Define a persistent data layer using  SQLite3 database

---

## Status

Work in progress. Currently only the `/healthcheck` endpoint is implemented.

---

