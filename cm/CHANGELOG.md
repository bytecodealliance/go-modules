# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [v0.3.0] — 2025-05-25

### Added

- Mutating methods `SetOK` and `SetErr` on `result` types (`Result[Shape, OK, Err]`).

### Changed

- Breaking: `BoolResult` is now represented as a `uint8` rather than a `bool`. This fixes an issue with TinyGo where `bool` values are treated distinctly from `uint8`. See [#344](https://github.com/bytecodealliance/go-modules/issues/344) for more information.

### Fixed

- [#344](https://github.com/bytecodealliance/go-modules/issues/344): the memory representation of `option` and `result` now use `uint8` instead of `bool` for the discriminator. LLVM optimizes `bool` values into a single bit, which breaks WIT variants where the associated types share memory.

## [v0.2.2] — 2025-03-16

### Fixed

- Updated error handling to address a [build issue](https://github.com/tinygo-org/tinygo/issues/4810) in TinyGo.

## [v0.2.1] — 2025-03-16

### Fixed

- Fixed cyclical dependency on package `encoding/json` when package `cm` is used in TinyGo package `syscall`. Files that import `encoding/json` will have a `_json.go` suffix and can be excluded when this package is copied into `std`.

## [v0.2.0] — 2025-03-15

### Added

- Initial support for Component Model [async](https://github.com/WebAssembly/component-model/blob/main/design/mvp/Async.md) types `stream`, `future`, and `error-context`.
- Initial support for JSON serialization of WIT `list`, `enum`, and `record` types.
- Added `cm.CaseUnmarshaler` helper for text and JSON unmarshaling of `enum` and `variant` types.

### Changed

- Breaking: package `cm`: removed `bool` from `Discriminant` type constraint. It was not used by code generation.

## [v0.1.0] — 2024-12-14

Initial version, extracted into module [`go.bytecodealliance.org/cm`](https://pkg.go.dev/go.bytecodealliance.org/cm).

[Unreleased]: <https://github.com/bytecodealliance/go-modules/compare/cm/v0.3.0..HEAD>
[v0.3.0]: <https://github.com/bytecodealliance/go-modules/compare/cm/v0.2.2..cm/v0.3.0>
[v0.2.2]: <https://github.com/bytecodealliance/go-modules/compare/cm/v0.2.1..cm/v0.2.2>
[v0.2.1]: <https://github.com/bytecodealliance/go-modules/compare/cm/v0.2.0..cm/v0.2.1>
[v0.2.0]: <https://github.com/bytecodealliance/go-modules/compare/cm/v0.1.0..cm/v0.2.0>
[v0.1.0]: <https://github.com/bytecodealliance/go-modules/tree/cm/v0.1.0>
