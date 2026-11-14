# HTTP API

This package will expose the scanner through HTTP in **Phase 3**. It will validate
requests, call the same scanning flow used by the CLI, and encode responses.

Planned contents:

- Router and server setup, with request timeouts and middleware.
- Request and response types, input validation, and error-to-HTTP-status mapping.
- Handlers for `POST /scans`, `GET /scans`, `GET /scans/{id}`,
  `GET /scans/{id}/findings`, and `GET /health` as their backing features become
  available.
- HTTP handler tests using Go's `net/http/httptest` package.

Keep scanning logic in [engine](../engine/) and security checks in
[rules](../rules/). The first API will run scans synchronously; durable scan
history follows with [storage](../storage/) in Phase 4, and queued scans follow
with [worker](../worker/) in Phase 5.

**Current phase:** planned only. Implement the Dockerfile parser before starting
HTTP handlers.
