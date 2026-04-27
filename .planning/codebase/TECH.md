# Technology Stack

## Language & Runtime
- Root module: Go 1.22.0 (toolchain go1.22.1)
- `gofp` sub-module: Go 1.18 (minimum for generics)

## Dependencies

### Root module (`github.com/msales/gox`)
No external dependencies. Pure Go standard library.

### `gofp` sub-module (`github.com/msales/gox/gofp`)
| Dependency | Purpose |
|---|---|
| `github.com/stretchr/testify` v1.8.0 | Test assertions (test-only) |

## Build System
- No Makefile present.
- `go build ./...` compiles the library.
- Dockerfile builds in `msales/go-builder:1.16-base-1.0.0` for CI validation (library only, no binary output).
- Multi-module structure requires building/testing each module separately.

## Testing
- `go test ./...` for root module packages.
- `cd gofp && go test ./...` for the gofp sub-module.
- Uses `testify/assert` throughout.
- Table-driven test style.
- Every package has corresponding `*_test.go` files.

## CI/CD
- Dockerfile serves as CI build check (compilation only).
- No `.github/workflows/` or other CI config found.
