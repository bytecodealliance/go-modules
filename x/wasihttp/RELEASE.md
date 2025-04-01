# Release Process

This document describes how to release a new version of the `wasihttp` package.

## Release Steps

1. Update the [CHANGELOG.md](CHANGELOG.md) with the changes since the last release.
2. Ensure all tests pass and examples work correctly.

## Creating a Release

Create a new git tag for the release:

```bash
git tag x/wasihttp/v0.1.0
git push origin x/wasihttp/v0.1.0
```

## Version Numbering

The `wasihttp` package follows [semantic versioning](https://semver.org/).
