<div align="center">
  <img src="docs/assets/images/godarajamascott.png" alt="go-daraja" width="100" />

  <h1>go-daraja</h1>

  A lightweight, idiomatic Go client for Safaricom’s Daraja (M-Pesa) APIs.

  [![Go](https://img.shields.io/badge/go-%3E%3D1.22-00ADD8?logo=go&logoColor=white)](https://go.dev/)
  [![License](https://img.shields.io/github/license/silaselisha/go-daraja)](LICENSE)
  [![Docs](https://img.shields.io/badge/docs-online-blue)](https://silaselisha.github.io/go-daraja/)
</div>

---

## Overview

go-daraja provides a clean, consistent interface to the core Daraja payment flows. It prioritizes safe defaults, straightforward configuration, and context-aware operations suitable for production use.

**Key features**

- Functional options for client construction
- Context support for cancellation and timeouts
- Pluggable `http.Client`
- 12-factor configuration (environment variables, optional `.env`)
- Embedded X.509 certificates (sandbox and production)
- Coverage of STK Push, B2C, C2B, B2B, and Business Express

## Installation

```bash
go get github.com/silaselisha/go-daraja
```

```go
import daraja "github.com/silaselisha/go-daraja/pkg/handler"

client, err := daraja.NewClient()
if err != nil {
    log.Fatal(err)
}
```

## Configuration

Configuration is loaded from the environment. An optional `.env` file may be provided via `WithEnvFile`.

| Variable | Required | Description |
|----------|----------|-------------|
| `MPESA_ENVIRONMENT` | No | `sandbox` (default) or `production` |
| `DARAJA_CONSUMER_KEY` | Yes | Consumer key |
| `DARAJA_CONSUMER_SECRET` | Yes | Consumer secret |
| `DARAJA_BUSINESS_SHORT_CODE` | Yes* | Business short code |
| `DARAJA_PASS_KEY` | Yes* | Pass key |
| `DARAJA_CALL_BACK_URL` | Yes* | Callback URL |
| `DARAJA_ACCOUNT_REF` | No | Account reference |
| `DARAJA_INITIATOR_NAME` | B2C/B2B | Initiator name |
| `DARAJA_INITIATOR_PASSWORD` | B2C/B2B | Initiator password |

\* Required for most payment operations.

```bash
export MPESA_ENVIRONMENT=sandbox
export DARAJA_CONSUMER_KEY=your_key
export DARAJA_CONSUMER_SECRET=your_secret
export DARAJA_BUSINESS_SHORT_CODE=174379
export DARAJA_PASS_KEY=your_pass_key
export DARAJA_CALL_BACK_URL=https://your.app/callback
```

## Supported APIs

- **STK Push** (Lipa na M-Pesa Online)
- **Business to Customer** (B2C)
- **Customer to Business** (C2B)
- **Business to Business** (B2B)
- **Business Express Checkout**
- **OAuth** access token generation

## Usage

### STK Push

```go
res, err := client.NIPush("payment for order #123", "0712345678", 100)
```

### Business to Customer

```go
res, err := client.BusinessToConsumer(
    100, 0, "0712345678", "salary",
    "https://example.com/timeout",
    "https://example.com/result",
)
```

### Customer to Business (register URLs)

```go
res, err := client.CustomerToBusiness(
    "https://example.com/confirm",
    "https://example.com/validate",
    1,
)
```

### Business Express Checkout

```go
res, err := client.BusinessExpressCheckout(
    "PAY-123",
    "https://example.com/callback",
    "YourBiz",
    "600000",
    100,
)
```

### Functional options

```go
client, err := daraja.NewClient(
    daraja.WithEnvFile("."),
    daraja.WithHTTPClient(customHTTPClient),
    daraja.WithLogger(logger),
)
```

Context-aware variants (`NIPushCtx`, etc.) are available for cancellation and timeouts.

## Project Structure

```
pkg/
  handler/          # Public client and operations
  internal/
    auth/           # Token generation
    builder/        # URL, phone, and timestamp helpers
    config/         # Configuration loader
    x509/           # Embedded certificates and security credentials
example/            # Sample usage
docs/               # Documentation site
```

## Development

Requires Go 1.22 or later.

```bash
go test ./...
go vet ./...
```

## Documentation

Full documentation is available at [silaselisha.github.io/go-daraja](https://silaselisha.github.io/go-daraja/).

## Contributing

Contributions are welcome. Please read [CONTRIBUTING.md](CONTRIBUTING.md) before opening a pull request.

## Security

Report security issues privately by following the process in [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
