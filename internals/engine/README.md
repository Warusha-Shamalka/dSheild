# Scan orchestration and policy evaluation

This package will connect file discovery, parsers, and rules into a reusable scan
flow. It is part of **Phase 1**, with recursive scanning and CLI integration added
as Phase 2 develops.

Planned contents:

- `scanner.go`: discover or accept input files, classify their format, invoke the
  appropriate parser, and aggregate findings and errors.
- `registry.go`: hold explicitly registered rules and select them by input format.
- `engine.go`: evaluate applicable rules against parsed resources and collect
  structured findings.
- Tests for rule selection, combined findings, and parse-error propagation.

The intended flow is:

```text
Input file -> parser -> structured model -> applicable rules -> findings
```

Use [parser](../parser/) to read syntax, [model](../model/) for shared data, and
[rules](../rules/) for individual checks. Register concrete rules in application
startup code under [cmd](../../cmd/); the scanner should not need edits whenever
a new rule is added.

**Current phase:** finish and test the Dockerfile parser first, then connect one
rule. Terminal formatting, HTTP responses, and database queries belong in their
own layers. A parse failure must remain distinguishable from a successful scan
with no findings.
