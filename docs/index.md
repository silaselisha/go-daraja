# Go Daraja (M‑Pesa) Client

A lightweight, idiomatic Go client for Safaricom's Daraja (M‑Pesa) APIs. Clean configuration, safe defaults, and a consistent API for core payment flows.

![](assets/images/godarajamascott.png){ width=120 }

## Highlights

- Simple, composable client construction with functional options
- Context-aware APIs for cancellation and timeouts
- Embedded X.509 certs and typed error model
- Covers STK Push, B2C, C2B, B2B and Business Express
- 12‑factor configuration from environment (optional `.env`)

## Quick links

- Getting started: [Installation](getting-started/installation.md) · [Quickstart](getting-started/quickstart.md)
- Usage guides: [STK Push](usage/ni-push.md), [B2C](usage/b2c.md), [C2B](usage/c2b.md), [B2B](usage/b2b.md), [Business Express](usage/business-express.md)
- Reference: [API](reference/api.md), [Environment variables](reference/environment.md)
- Support: [Troubleshooting](support/troubleshooting.md) · [FAQ](support/faq.md)

## Project structure

```text
pkg/
  handler/          # Public client and operations
  internal/
    auth/           # Basic auth token creation
    builder/        # Helpers (base URLs, phone formatting, timestamp)
    config/         # Viper-based configuration loader
    x509/           # go:embed certs + security credential generation
example/            # Example .env and sample usage
```
