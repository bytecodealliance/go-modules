# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

### Fixed

- Breaking: fixed a cyclical dependency introduced in [v0.2.0] between package `encoding/json` and TinyGo `syscall`. To enable JSON support for Component Model `list` and other types, programs must now explicitly import package `cm/json`:
	`import _ "go.bytecodealliance.org/cm/json"`

## [v0.2.0] — 2025-03-15

### Added

- Initial support for Component Model [async](https://github.com/WebAssembly/component-model/blob/main/design/mvp/Async.md) types `stream`, `future`, and `error-context`.
- Initial support for JSON serialization of WIT `list`, `enum`, and `record` types.
- Added `cm.CaseUnmarshaler` helper for text and JSON unmarshaling of `enum` and `variant` types.

### Changed

- Breaking: package `cm`: removed `bool` from `Discriminant` type constraint. It was not used by code generation.

## [v0.1.0] — 2024-12-14

Initial version, extracted into module [`go.bytecodealliance.org/cm`](https://pkg.go.dev/go.bytecodealliance.org/cm).

[Unreleased]: <https://github.com/bytecodealliance/go-modules/compare/cm/v0.2.0..HEAD>
[v0.2.0]: <https://github.com/bytecodealliance/go-modules/compare/cm/v0.1.0..cm/v0.2.0>
[v0.1.0]: <https://github.com/bytecodealliance/go-modules/tree/cm/v0.1.0>
