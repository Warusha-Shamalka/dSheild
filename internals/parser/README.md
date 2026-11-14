# Configuration parsers

This package will turn infrastructure configuration files into structured data
that security rules can inspect. This is the starting point for **Phase 1: Core
Static Scanner** in the [roadmap](../../docs/DeployShield_RoadmapV2.pdf), printed
page 6. The roadmap PDF is currently ignored by Git, so this link requires the
local copy.

The chosen learning order is **Dockerfile -> Kubernetes -> Terraform**. Finish a
small, tested Dockerfile parser before adding the other formats.

## Planned contents

- `dockerfile.go`: parse Dockerfile input and map the result into dSheild's model.
- `dockerfile_test.go`: table-driven tests for instructions, locations, and errors.
- `kubernetes.go` and its tests later: parse YAML, including multiple documents,
  while retaining useful source positions.
- `terraform.go` and its tests later: parse HCL and preserve unresolved expressions
  instead of pretending every value is a known literal.

Use [model](../model/) for shared parsed data and [tests](../../tests/) for fixture
files. The filenames above are proposed; no parser is implemented yet.

## Build the first Dockerfile parser step by step

1. **Define the output.** Start with a Dockerfile containing a file path and an
   ordered list of instructions. Each instruction needs its keyword, original
   text, arguments/flags, and one-based start/end lines. Keep enough structure to
   distinguish JSON-form arguments from shell-form text.
2. **Create one tiny fixture.** Begin with `FROM alpine:3.22`, `RUN echo hello`,
   and `USER 1000` on separate lines. The image reference is just sample text;
   parsing it does not pull or build the image.
3. **Parse syntax and map it.** The roadmap proposes a small tokenizer. A practical
   implementation choice is to wrap an established Dockerfile syntax parser,
   such as BuildKit's `github.com/moby/buildkit/frontend/dockerfile/parser`, and
   build dSheild's input handling, model mapping, and diagnostics around it.
   This avoids treating simple line splitting as a complete Dockerfile grammar.
   Select and add the dependency when implementation starts.
4. **Preserve source locations.** A multiline instruction is one logical
   instruction spanning several physical lines. Normalize the keyword for checks
   while keeping the original spelling and source text available.
5. **Test realistic syntax.** Cover comments, blank lines, CRLF endings, line
   continuations, parser directives such as `# escape=`, JSON-form instructions,
   instruction flags, and multistage builds. Record which stage owns an
   instruction; a `FROM` can reference an earlier stage or `scratch` as well as an
   external image. Add heredoc coverage before claiming broad Dockerfile support.
6. **Handle failures explicitly.** Return errors for unreadable input and syntax
   failures detected by the parser, with the file path and line when available.
   Syntax parsing alone does not validate every Docker build constraint. Document
   any unsupported forms; never turn a failure into a successful empty result.
7. **Verify the output.** Assert instruction order, arguments, stage information,
   and exact source lines in tests. Once implemented, run
   `go test ./internals/parser` from the repository root. This command has no Go
   package to test until the source and test files exist.

For the first fixture, the expected starting point is:

| Keyword | Arguments | Start line | End line |
| --- | --- | --- | --- |
| FROM | alpine:3.22 | 1 | 1 |
| RUN | echo hello | 2 | 2 |
| USER | 1000 | 3 | 3 |

The first milestone is complete when the parser returns those structured
instructions and tests verify them. Then extend the syntax fixtures before
connecting [rules](../rules/) and [engine](../engine/).

## Keep the parser focused

Parse input as data. Do not execute `RUN`, build images, resolve remote image
metadata, or require a Docker daemon. Leave security decisions to rules and output
formatting to callers.

A missing `USER` instruction does not prove that an image runs as root: the base
image may define its own user. Preserve the available evidence so later rules can
report that distinction accurately.

## Repository naming note

The checkout currently uses `internals/`; the roadmap suggests `internal/`.
Only the exact directory name `internal` activates Go's internal-package import
restriction. The current spelling can contain Go packages, but does not provide
that restriction. These READMEs preserve the existing layout.
