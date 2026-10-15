# Contributing to go-daraja

Thank you for your interest in contributing. This guide outlines the process for proposing changes.

## Getting Started

1. Fork the repository and create a branch from `master`.
2. Make your changes.
3. Ensure tests pass:

   ```bash
   go test ./...
   ```

4. Open a pull request against `master`.

## Pull Request Guidelines

- Keep changes focused. A pull request should address a single concern.
- Provide a clear description of the change and the reason for it. Link related issues when applicable.
- Prefer small, reviewable commits with descriptive messages.
- Keep the branch up to date with `master` to avoid merge conflicts.
- All tests must pass before a change can be merged.

For larger or speculative changes, open an issue for discussion first.

## Code Style

Follow standard Go conventions. Run `go vet ./...` and format code with `gofmt` before submitting.

## Reporting Issues

Use the issue templates when reporting bugs or requesting features. Include reproduction steps, expected behavior, and relevant environment details.
