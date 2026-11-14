# Shared scan data

This package will contain the data structures passed between parsers, rules, the
engine, and eventually the CLI, API, and storage layer. Keep these types independent
of HTTP handlers, database drivers, and third-party parser internals.

Planned contents:

- `dockerfile.go`: the initial Dockerfile representation and its instructions.
  Preserve the file path, instruction keyword, arguments and flags without losing
  meaningful syntax, original instruction text, and one-based start/end lines.
  Add build-stage information as multistage support is implemented.
- `resource.go`: file-type identifiers and shared resource fields when another
  input format makes them useful. Keep format-specific details rather than forcing
  Dockerfile instructions into Kubernetes or Terraform structures.
- `finding.go`: rule ID, severity, file path, source location where available,
  message, and remediation recommendation.
- `severity.go`: `CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, and `INFO` values, with an
  explicit ordering for comparisons.
- `scan.go`: scan metadata, timestamps, findings, and status when the scan flow
  needs them. The roadmap requires a scan timestamp in reported findings; retain
  that association when defining report and storage types.

**Current phase:** start with only the Dockerfile and instruction types required
by [parser](../parser/). Add findings and severities when implementing the first
rule, and scan metadata when introducing orchestration. The parser should produce
structured input, not security findings.
