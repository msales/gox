@~/CLAUDE.md

## Project Description

**Language:** Go 1.22 (root module), Go 1.18 (`gofp` sub-module)  
**Module:** `github.com/msales/gox`  
**Purpose:** A collection of general-purpose Go utility packages providing slice operations, functional programming helpers (with generics), math utilities, networking helpers, GDPR data protection, random probability functions, and a thread-safe FIFO queue.  
**Key dependencies:**
- Root module: no external dependencies (zero deps)
- `gofp` sub-module: `github.com/stretchr/testify` (testing only)

## Directory Structure

```
gdpr/           GDPR data protection: IP anonymization, device ID masking
gofp/           Functional programming with generics: Map, Filter, Any, Contains, Diff, Union, Intersection, Unique, Set, IndexedSet, Reduce, Reindex, Merge, Count
gofp/setx/      Set algebra: Sum (OR) and Product (AND) condition merging with operators (In, NotIn, All, None)
mathx/          Math utilities: rounding to N decimal places
netx/           Network utilities: IP type wrapper, raw IP to uint32 conversion
queuex/         Thread-safe generic FIFO queue
randx/          Probability utilities: Happens, WhichHappens
slicex/         Slice type conversion and utilities (string, int, float, uint variants + reflection-based Contains/Filter)
```

## Build, Test & Lint

```bash
# Build
go build ./...

# Run tests (root module)
go test ./...

# Run tests (gofp sub-module)
cd gofp && go test ./...

# Run tests with verbose output
go test -v ./...

# Vet
go vet ./...

# Docker build (CI validation only)
docker build --build-arg GITHUB_TOKEN=$GITHUB_TOKEN .
```

No Makefile is present. The Dockerfile is used solely for CI build validation.

## Code Style

### Imports

Three groups separated by blank lines:
1. Standard library
2. Third-party packages
3. Internal/msales packages

### Naming Conventions

- **Packages:** short, descriptive, lowercase with `x` suffix convention for extended stdlib equivalents (`slicex`, `mathx`, `netx`, `queuex`, `randx`)
- **Interfaces:** type constraint interfaces using generics (e.g., `IP interface { net.IP | netx.IP | uint32 | string }`)
- **Exported/unexported:** all utility functions exported; helper functions unexported
- **Generics:** used extensively in `gofp` and `queuex` packages (Go 1.18+ features)

### Error Handling

- Functions that can fail return `(result, error)` tuples (e.g., `StringToInt`).
- Functions with impossible failure modes return results directly (e.g., `IntToString`).
- Reflection-based functions panic on type mismatch (e.g., `slicex.Contains`).

### Testing

- `testify/assert` for assertions
- Table-driven tests throughout
- Every package has comprehensive `*_test.go` files with edge cases
- `gofp` sub-module has its own `go.mod` and test suite

## Branching & Git

- **Never** work directly on `main` or `master`. Always create a new branch off the latest `main`/`master`.
- Before creating a branch, **always pull** the latest changes from the remote (`git pull origin master`).
- If the user provides a task identifier in the format `TRK-XXXX`, use it as the branch name (e.g., `TRK-1234`). Otherwise, **ask the user** for a task ID or branch name before proceeding.
- **Never push** to `main` or `master` directly.
- Commit changes after each meaningful iteration. Use your judgment to decide when progress should be saved -- prefer smaller, atomic commits over large monolithic ones.
- Write clear, concise commit messages. Use conventional commit format when appropriate (e.g., `feat:`, `fix:`, `refactor:`, `test:`).
