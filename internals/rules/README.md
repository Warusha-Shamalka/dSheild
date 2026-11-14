# Security rules

This package will contain independent security checks against parsed resources.
It follows the roadmap's **Phase 1** rule architecture: metadata plus an evaluation
method, selected through a registry rather than hard-coded into the scanner.

Planned contents:

- `rule.go`: the small rule contract and metadata such as ID, description,
  severity, applicable format, recommendation, and reference.
- `dockerfile/`: Dockerfile checks, introduced first.
- `kubernetes/` and `terraform/`: format-specific checks introduced later.
- Unit tests next to each implementation, with focused clean and insecure
  examples from [tests](../../tests/).

After [parser](../parser/) works, start with one explicit-root check for
`USER root` or `USER 0`. Define whether it evaluates every stage or only the final
runtime stage, and test that scope. A later `USER` can change the effective user
within a stage, so an earlier instruction alone does not establish its final user.

The roadmap's illustrative claim that a missing `USER` always means root is too
strong: a base image may already define a non-root user. Report only what the
available configuration establishes.

Later checks can inspect external base-image references and secret-like values.
An earlier stage alias or `scratch` in `FROM` needs different handling from an
external image reference.

Rules return [findings](../model/); they do not parse files, print output, or write
to PostgreSQL. Register them explicitly at startup and let
[engine](../engine/) select and evaluate them. **Current phase:** planned only;
the parser is the first task.
