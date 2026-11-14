# Persistent scan history

This package will store and retrieve scan data in PostgreSQL in **Phase 4**.
Persistence is introduced after the scanner and synchronous API work.

Planned contents:

- Database connection setup, configuration validation, and clean shutdown.
- Repository methods and SQL for repositories, scans, findings, and rule metadata.
- Transactions for operations that must succeed or fail together, and explicit
  handling of missing records and database failures.
- Integration tests against a dedicated test database.

Versioned schema changes belong in [migrations](../../migrations/). Shared domain
data belongs in [model](../model/); HTTP request handling belongs in
[api](../api/). Load credentials from runtime configuration rather than embedding
them in source or fixtures.

**Current phase:** planned only. The first Dockerfile parser and terminal scanner
can work entirely without PostgreSQL.
