# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- WIT labels with uppercase acronyms or [initialisms](https://go.dev/wiki/CodeReviewComments#initialisms) are now preserved in Go form. For example, the WIT name `time-EOD` will result in the Go name `TimeEOD`.
- `OK` is now a predefined initialism.

## [v0.1.0] — 2024-07-04

Initial version, supporting [TinyGo](https://tinygo.org/) + [WASI](https://wasi.dev) 0.2 (WASI Preview 2).

### Known Issues
- [#95](https://github.com/ydnar/wasm-tools-go/issues/95): `variant` and `result` types without fully-packed `data` shape types will not correctly represent all associated types.
- [#111](https://github.com/ydnar/wasm-tools-go/issues/111): `flags` types with > 32 labels are not correctly supported. See [component-model#370](https://github.com/WebAssembly/component-model/issues/370) and [wasm-tools#1635](https://github.com/bytecodealliance/wasm-tools/pull/1635) for more information.
- [#118](https://github.com/ydnar/wasm-tools-go/issues/118): Canonial ABI [post-return](https://github.com/WebAssembly/component-model/blob/main/design/mvp/CanonicalABI.md#canon-lift) functions to clean up allocations are not currently generated.
- Because Go does not have a native tagged union type, pointers represented in `variant` and `result` types may not be visible to the garbage collector and may be freed while still in use.
- Support for mainline [Go](https://go.dev/).

[Unreleased]: <https://github.com/ydnar/wasm-tools-go/compare/v0.1.0...HEAD>
[v0.1.0]: <https://github.com/ydnar/wasm-tools-go/tree/v0.1.0>