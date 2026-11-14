# Scanner fixtures and integration tests

This folder will contain input files and tests that exercise several scanner
packages together. Keep ordinary Go unit tests beside the package they test;
for example, parser unit tests belong in [parser](../internals/parser/).

Planned contents:

```text
tests/
  fixtures/
    dockerfiles/             # First: syntax examples, then clean/insecure cases
    kubernetes/              # Later: YAML examples and rule cases
    terraform/               # Later: HCL examples and rule cases
  scan_integration_test.go   # Parser -> rules -> findings
```

These paths are proposed and will be created as their tests are implemented.

## First parser fixtures

Start with a three-instruction Dockerfile containing `FROM`, `RUN`, and `USER`.
Then add focused fixtures for:

- Comments, blank lines, mixed-case keywords, and CRLF line endings.
- Backslash continuations and exact start/end line numbers.
- Parser directives, JSON-form arguments, and instruction flags.
- Multiple build stages, stage aliases, and `FROM scratch`.
- Malformed syntax and the diagnostics it produces.
- Heredocs before claiming broad syntax support.

Keep expected instruction order, arguments, and locations explicit in table-driven
parser tests. Use Go's temporary-directory helpers for missing-file and permission
cases where practical, rather than committing machine-specific paths.

## Later rule and scan tests

Pair each rule's insecure fixture with a clean one. Ideally, each insecure example
triggers only the check it is designed to test. Add a named regression fixture
whenever a parser bug or false positive is fixed. Check that malformed input is
reported as an error rather than a clean scan.

Run `go test ./...` from the repository root once tests exist. Static scanner
fixtures are read as text: they do not require Docker, Kubernetes, or a cloud
account. Database and queue integration tests become relevant in later phases.
