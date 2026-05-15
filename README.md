# dive

A fast, minimal developer CLI for accessing Docker containers and database CLIs.

## Philosophy

- Fast, minimal, keyboard-first, and clean
- No magic, no overengineering, no global state
- Inspired by lazydocker, lazygit, kubectl, gh CLI, pnpm

## Installation

```
go install github.com/mzulfanw/dive/cmd/dive@latest
```

## Usage

```
dive ps           # Show running containers
dive pg           # Enter psql in PostgreSQL container
dive redis        # Enter redis-cli in Redis container
dive mysql        # Enter mysql client in MySQL container
dive sh api       # Enter shell in 'api' container
dive logs api     # Stream logs from 'api' container
```

## Examples

- `dive pg` → Connects to the first running PostgreSQL container
- `dive sh web` → Opens a shell in the 'web' container
- `dive logs db` → Streams logs from the 'db' container

## Architecture Overview

- **cmd/**: CLI entrypoint
- **internal/cli/**: Cobra commands
- **internal/docker/**: Docker shell logic
- **internal/resolver/**: Container detection
- **internal/runtime/**: Runtime abstraction
- **internal/types/**: Data types
- **internal/utils/**: Helpers
- **pkg/output/**: Output formatting

## Roadmap

- TUI mode
- Kubernetes/Podman support
- Docker SDK migration
- More database CLIs

## License

MIT
