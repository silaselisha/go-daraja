# Changelog

All notable changes to this project are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).

## [Unreleased]

### Changed
- **Breaking:** failed requests now return a non-nil Go `error` (`*handler.APIError`) for network failures, non-2xx responses, M-Pesa API `errorCode` payloads and non-JSON bodies. The response is still returned with `ErrorCode` and `ErrorMessage` populated. ([#53](https://github.com/silaselisha/pokea/issues/53))
- **Breaking:** `NewClient`, `NewPokeaClient` and `NewDarajaClient` now return an error when token acquisition fails, instead of building a client with an empty access token.
- `NewClient` uses the injected HTTP client (`WithHTTPClient`) for token acquisition.
- Renamed Daraja identifiers to Pokea: `NewDarajaClient` → `NewPokeaClient`, `Daraja*` types → `Pokea*`, `DARAJA_*` environment variables → `POKEA_*`. The embedded response field is now `PokeaErrorParams`.

### Added
- `handler.ErrUnreachable` for `errors.Is` checks, and `Failed()` / `AsError()` helpers on responses.

### Deprecated
- `NewDarajaClient` and the `Daraja`, `DarajaClient`, `DarajaResParams`, `DarajaErrorParams` and `DarajaAuth` aliases. They will be removed in the next major release.
- `DARAJA_*` environment variables. They are still read when the matching `POKEA_*` variable is unset.

### Removed
- `AUTHORS` and `CONTRIBUTORS` files.

### Fixed
- Configuration loading now reads OS environment variables when no `.env` file is present. Previously it failed with a missing-credentials error.
- A failure to close a response body no longer terminates the process via `log.Fatalf`.

## [0.9.0] - 2025-10-26

### History rewrite
- The `v0.9.0` tag was moved to rewritten history. Nine commits originally authored by an automated agent identity are now attributed to the repository owner. File contents at the tag are byte-for-byte identical; only commit metadata and hashes changed.
- If you fetch `github.com/silaselisha/go-daraja@v0.9.0` directly from GitHub (for example with `GOPRIVATE` or `GONOSUMDB`), you may see a checksum mismatch. Fetching through the Go module proxy is unaffected.

[Unreleased]: https://github.com/silaselisha/pokea/compare/v0.9.0...HEAD
[0.9.0]: https://github.com/silaselisha/pokea/releases/tag/v0.9.0
