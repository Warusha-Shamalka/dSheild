# Background scan worker

This package will process queued scans in **Phase 5**, after synchronous scans and
PostgreSQL history work. The roadmap uses a Redis-backed queue at this stage.

Planned contents:

- Job payload types and queue producer/consumer integration.
- A worker loop with cancellation, timeouts, and graceful shutdown.
- Scan state transitions: `QUEUED -> RUNNING -> DONE/FAILED`.
- Calls to the existing [engine](../engine/) and [storage](../storage/) packages.
- Bounded retries for transient failures, with tests for failures and duplicate
  delivery so a retried job does not corrupt scan history.

Keep the worker as another caller of the scanner; it should not implement a
second parser or rules engine. Do not execute code from scanned repositories.

**Current phase:** planned only. Queues, Redis, and concurrency are unnecessary
for building the first Dockerfile parser.
