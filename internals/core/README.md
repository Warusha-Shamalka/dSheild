# Shared application support

This package will hold the small pieces of application support needed across
dSheild. In the roadmap's suggested layout, `core` means configuration, logging,
and common errors; it is not a second scanner implementation.

Planned contents:

- `config.go`: configuration loading and validation when configuration is needed.
- `logging.go`: setup for Go's structured `log/slog` logger.
- `errors.go`: shared error categories only when several packages need to
  distinguish them. Wrap errors with file paths and operation context.
- Focused unit tests next to the code they exercise.

Scan orchestration belongs in [engine](../engine/), and shared scan data belongs
in [model](../model/). Avoid turning this package into a collection of unrelated
helpers.

**Current phase:** add only what the parser actually needs. Ordinary wrapped Go
errors are enough to begin; an extensive configuration or error framework is not
required.
