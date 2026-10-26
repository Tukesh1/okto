# Okto

Okto is a terminal-first GitHub intelligence CLI. The planned MVP provides a
repository overview, a responsive terminal dashboard, local SQLite caching,
historical snapshots, basic charts, repository comparison, and JSON output for
scripting.

## Status

This repository is at **Phase 0: scaffold**. The module, package boundaries, CLI
contract, build tooling, and CI exist. GitHub fetching, SQLite persistence,
analytics, charts, and the Bubble Tea dashboard are **not implemented yet**.

`okto --help` and `okto --version` work. `okto <owner/repo>` and
`okto compare <owner/repo> <owner/repo>` currently return a clear
"not implemented" error.

## Requirements

- Go 1.27 or newer
- `git`
- Optional: `gh` for the planned GitHub authentication fallback

## Build and test

```bash
make build
./bin/okto --help

make test
make check
```

`make check` verifies formatting, runs `go vet ./...`, and runs
`go test -race ./...`.

## Planned commands

```bash
okto owner/repo
okto owner/repo --json
okto owner/repo --refresh
okto compare owner/one owner/two
okto compare owner/one owner/two --json
```

When stdout is a TTY, the default output is the interactive dashboard. When
stdout is piped, the default output is a plain-text summary. `--json` always
produces machine-readable JSON on stdout and sends diagnostics to stderr.

See [docs/contracts.md](docs/contracts.md) for the planned CLI, exit-code, and
JSON contracts.

## Planned architecture

```text
cmd/okto
internal/cli        argument parsing and exit codes
internal/app        cached-first orchestration and refresh
internal/github     REST client, auth, ETags, rate limits
internal/store      SQLite schema, migrations, cache queries
internal/analytics  pure metric, trend, and comparison logic
internal/model      shared domain and view types
internal/chart      pure terminal chart rendering
internal/output     JSON schema v1 and plain text
internal/tui        responsive Bubble Tea dashboard
internal/config     paths, environment, and auth policy
```

## Planned authentication and storage

- Tokens are read from `GITHUB_TOKEN`, then `GH_TOKEN`, then `gh auth token`.
- There is no `--token` flag.
- Tokens are never stored in SQLite or printed in errors.
- Historical snapshots live in a local SQLite database. The data directory will
  be overridable with `OKTO_DATA_DIR` during development.
- No telemetry is planned.

## Development rules

Project instructions are loaded from `.opencode/skills/okto-dev/AGENTS.md`.
The Go module and binary use **Okto**.
