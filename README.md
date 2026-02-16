# memo

A Go-based CLI for writing ideas down immediately.

## Core Policy

- Single binary written in Go
- One-step memo creation with `memo <description>`
- File name format: `yyyymmdd_<description>.md`

## Usage

### 1) Initialize

```bash
memo init
```

Configure memo directory interactively. If left empty, current working directory is used.
Configure memo directory interactively. If left empty, XDG data path is used:
`$XDG_DATA_HOME/memo` (fallback: `~/.local/share/memo`).

### 2) Create a memo

```bash
memo buy milk
```

Example output file: `20260216_buy_milk.md`

### 3) Search memos

```bash
memo search bmk
```

Performs simple fuzzy matching (in-order character match) against memo file names.

### 4) Open a memo

```bash
memo open bmk
```

Finds the best fuzzy match and opens it in `$EDITOR` (`vi` fallback).

### 5) Show revision

```bash
memo version
```

Prints the embedded revision value.

### 6) List recent memos

```bash
memo list --today
memo list --week
```

Prints memo file paths created today or within the last 7 days.

## Development

```bash
mise run all
```

## Tooling

- `goreleaser`: release artifact generation/signing/SBOM (`.goreleaser.yaml`)
- `lefthook`: local Git hooks (`lefthook.yml`)
- `renovate`: dependency update automation (`.github/renovate.json`)
- `golangci-lint`: linting (`.golangci.yml`)
