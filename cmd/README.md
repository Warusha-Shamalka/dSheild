# Application entry points

This folder will contain the executable programs that start dSheild. Keep each
entry point small: read arguments or configuration, connect the required packages,
and call the scanner or service.

Planned contents:

- `dshield/main.go`: the command-line entry point, including the future `scan`
  command, help text, output options, and exit codes.
- API and worker entry points when those services are introduced. Separate
  executables are only needed when they must run independently.

The roadmap introduces the full CLI in **Phase 2**. After the first parser works,
a minimal command can help demonstrate it. Parsing belongs in
[parser](../internals/parser/); scan orchestration belongs in
[engine](../internals/engine/).

The planned CLI exit codes are `0` for a completed scan below the configured
failure threshold, `1` for findings that meet that threshold, and `2` for an input
or scanner error. These commands and behaviors are not implemented yet.
