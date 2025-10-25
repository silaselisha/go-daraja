<div align="center" style="margin-bottom: 0px!important; padding: 0px;">
    <img src="./public/images/godarajamascott.png" alt="godaraja logo" height="100px"/>
</div>

<div style="align-items: center; margin-top: 0px !important; margin-bottom: 14px;" align="center">
    <p style="text-align: center;" align="center">
        <img src="https://img.shields.io/badge/go-%3E%3D1.22-blue?logo=go" />
        <img src="https://img.shields.io/github/license/silaselisha/go-daraja" />
        <img src="https://img.shields.io/github/actions/workflow/status/silaselisha/go-daraja/ci.yml?label=CI" />
    </p>
    <h1 style="font-size: 44px; font-weight: 800; padding: 0px;">go-daraja</h1>
</div>

## Overview

go-daraja is a lightweight, idiomatic Go client for Safaricom's Daraja (M-Pesa) APIs. It focuses on a clean developer experience, easy configuration across environments, and safe defaults for production.

### Features

- **Simple client**: One constructor with functional options
- **12‑factor config**: OS env first, optional `.env`, sane defaults
- **Context-aware**: `*_Ctx` variants for timeouts/cancellation
- **Pluggable HTTP**: Inject your own `http.Client`
- **Embedded certs**: Uses `go:embed` for sandbox/prod X509 certs
- **Consistent errors**: Typed error model internally, backward-compatible responses
- **Coverage of core flows**: STK Push, B2C, C2B, B2B, Business Express

## Installation

```bash
go get github.com/silaselisha/go-daraja
```

## Quickstart

Set required envs (12‑factor friendly). `.env` is optional and supported.

```bash
export MPESA_ENVIRONMENT=sandbox
export DARAJA_CONSUMER_KEY=your_key
export DARAJA_CONSUMER_SECRET=your_secret
export DARAJA_BUSINESS_SHORT_CODE=174379
export DARAJA_PASS_KEY=your_pass_key
export DARAJA_CALL_BACK_URL=https://your.app/callback
export DARAJA_ACCOUNT_REF=YourApp
```

```go
import (
    "log"
    daraja "github.com/silaselisha/go-daraja/pkg/handler"
)

func main() {
    client, err := daraja.NewClient()
    if err != nil { log.Fatal(err) }

    res, err := client.NIPush("test STK push", "0708374149", 1)
    if err != nil { log.Fatal(err) }
    log.Printf("%+v\n", res)
}
```

Backward compatible:

```go
client, err := daraja.NewDarajaClient(".")
```

## Configuration

Key environment variables:

- **MPESA_ENVIRONMENT**: `sandbox` | `production` (default: `sandbox`)
- **DARAJA_CONSUMER_KEY**, **DARAJA_CONSUMER_SECRET**: required
- **DARAJA_BUSINESS_SHORT_CODE**: required for STK/C2B
- **DARAJA_PASS_KEY**: required for STK
- **DARAJA_CALL_BACK_URL**: callback for STK
- **DARAJA_ACCOUNT_REF**: default reference label
- **DARAJA_INITIATOR_NAME**, **DARAJA_INITIATOR_PASSWORD**: required for B2C/B2B

You can also load from a `.env` file by placing it in your project root or by using `WithEnvFile("path")`.

## Usage

STK Push (NI):

```go
res, err := client.NIPush("payment for order #123", "0712345678", 100)
```

B2C:

```go
res, err := client.BusinessToConsumer(100, 0, "0712345678", "salary", "https://example.com/timeout", "https://example.com/result")
```

C2B register URLs:

```go
res, err := client.CustomerToBusiness("https://example.com/confirm", "https://example.com/validate", 1)
```

Business Express Checkout:

```go
res, err := client.BusinessExpressCheckout("PAY-123", "https://example.com/callback", "YourBiz", "600000", 100)
```

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

## Development

Prereqs: Go >= 1.22

- Run tests:
```bash
go test ./...
```

- Lint/inspect locally:
```bash
go vet ./...
```

## Contributing

Contributions are welcome! Please read `CONTRIBUTING.md` and open an issue or PR.

## Security

If you discover a security vulnerability, please follow the process in `SECURITY.md`.

## License

This project is licensed under the MIT License. See `LICENSE` for details.

## Stats

<p align="center">
    <img src="https://repobeats.axiom.co/api/embed/36b264b4be024052073f9c5703b102cd24693c62.svg" alt="go-daraja stats" title="Repobeats analytics image"/>
</p>
