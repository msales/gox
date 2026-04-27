# Architecture

## Overview
gox is a utility library organized as independent packages, each providing a focused set of helper functions. There is no service architecture -- it is a collection of pure functions and data structures used by other msales projects.

## Package Structure
| Package | Responsibility |
|---|---|
| `gdpr` | GDPR compliance: `ProtectIP` (anonymize IPv4/IPv6), `ProtectDeviceID` (mask device identifiers) |
| `gofp` | Functional programming with generics: `Map`, `MapMultiple`, `Filter`, `Any`, `ContainsAll`, `ContainsAtLeastOne`, `Diff`, `DiffWithKeyBuilder`, `Union`, `Intersection`, `Unique`, `Set`, `IndexedSet`, `Reduce`, `Reindex`, `Merge`, `Count` |
| `gofp/setx` | Set algebra for condition merging: `Sum` (OR logic), `Product` (AND logic) with `In`/`NotIn`/`All`/`None` operators |
| `mathx` | Math utilities: `Round` (round to N decimal places) |
| `netx` | Network helpers: `IP` type (wraps `net.IP`), `RawIPToUint`, `UintToIP` |
| `queuex` | Thread-safe generic `FIFO[T]` queue with `Push`, `Pop`, `Len` |
| `randx` | Probability functions: `Happens` (single probability), `WhichHappens` (weighted selection) |
| `slicex` | Slice type conversions (string, int, int32, int64, uint, uint32, uint64, float32, float64, interface) and generic `Contains`, `Filter`, `ContainsAtLeastOne` |

## Key Interfaces
- **`gdpr.IP`** -- type constraint: `net.IP | netx.IP | uint32 | string` (used with generics for `ProtectIP`)
- **`setx.Result[Value]`** -- generic struct for set algebra results with `Operator` and `Values`
- **`setx.Sum[Value]`** / **`setx.Product[Value]`** -- condition merger types with `Merge(operator, values...)` method

## Data Flow
Not applicable -- this is a stateless utility library. Each function takes inputs and returns outputs without side effects (except `randx` which uses `math/rand` global state).

## External Dependencies
- No external services.
- `gdpr` package depends on `netx` package within the same module.
- `gofp/setx` depends on `gofp` functions (same sub-module).
