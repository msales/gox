# Concerns & Risks

## Technical Debt
- **Mixed paradigms in `slicex`:** older reflection-based functions (`Contains`, `Filter` using `reflect`) coexist with newer generic functions (`ContainsAtLeastOne`). The reflection-based functions could be replaced with generics.
- **Duplicated `ContainsAtLeastOne`:** exists in both `slicex` and `gofp` packages with identical implementations.
- **Dockerfile references Go 1.16** but root module requires Go 1.22 -- the Dockerfile may fail to build.
- **`gofp` sub-module uses Go 1.18** while root module uses Go 1.22 -- version mismatch across the project.
- **`slicex` type-specific conversion functions** (IntToString, StringToInt, etc.) are repetitive; could be consolidated with generics.

## Missing Coverage
- No benchmarks for any package.
- `randx` tests may be non-deterministic (probability-based).
- `setx.Product.mergeAll` has a TODO comment: "implement when needed."
- No integration tests.

## Security
- `gdpr.ProtectIP` handles PII anonymization -- critical for GDPR compliance. Implementation looks correct (zeroes last octet for IPv4, last 4 octets for IPv6).
- `gdpr.ProtectDeviceID` masks last 2 characters -- adequate for basic anonymization.
- No secrets or credentials in the codebase.

## Performance
- `slicex` reflection-based functions have overhead from `reflect.ValueOf` -- consider migrating to generics.
- `gofp.ContainsAll` and `ContainsAtLeastOne` optimize for single-element case to avoid map allocation.
- `queuex.FIFO.Pop` copies the slice on each pop (slice reslicing) -- could accumulate memory for large queues that are not garbage collected.
- `randx` uses global `math/rand` source with implicit mutex contention under high concurrency.
