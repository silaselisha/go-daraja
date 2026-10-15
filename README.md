<p align="center">
  <img src="public/images/godarajamascott.png" width="120" alt="go-daraja">
</p>
<div align="center" style="line-height: 1;">
  <a href="https://go.dev/" target="_blank"><img alt="Go" src="https://img.shields.io/badge/go-%E2%89%A51.22-00ADD8?logo=go&logoColor=white"/></a>
  <a href="LICENSE" target="_blank"><img alt="License" src="https://img.shields.io/github/license/silaselisha/go-daraja"/></a>
  <a href="https://github.com/silaselisha/go-daraja/actions" target="_blank"><img alt="CI" src="https://img.shields.io/github/actions/workflow/status/silaselisha/go-daraja/ci.yml?label=CI"/></a>
  <a href="https://pkg.go.dev/github.com/silaselisha/go-daraja" target="_blank"><img alt="Go Reference" src="https://pkg.go.dev/badge/github.com/silaselisha/go-daraja.svg"/></a>
</div>
<br>
<div align="center">
  A lightweight, idiomatic Go client for Safaricom’s Daraja (M-Pesa) APIs.
</div>
<br>
---
# go-daraja: Go Client for Safaricom Daraja (M-Pesa)
## News
<!-- news:start -->
- [2026-10] **go-daraja** refined with a cleaner public surface, functional options, context-aware methods, and a GitHub Actions CI workflow.
- [2024-03] Initial public release covering STK Push, B2C, C2B, B2B, and Business Express.
<!-- news:end -->
<div align="center">
🚀 [Overview](#overview) | ⚡ [Installation](#installation) | 📦 [Usage](#usage) | 🔧 [Configuration](#configuration) | 📁 [Structure](#project-structure)
</div>
## Overview
go-daraja is a focused Go library for integrating with Safaricom’s Daraja (M-Pesa) APIs. It provides a clean, consistent interface to the core payment flows with safe defaults, 12-factor configuration, and context support for production use.
**Highlights**
- Functional options for client construction
- Context-aware methods for cancellation and timeouts
- Pluggable `http.Client`
- Embedded X.509 certificates (sandbox and production)
- Typed responses with backward-compatible error shapes
- Coverage of STK Push, B2C, C2B, B2B, and Business Express
## Installation
Requires Go 1.22 or later.
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
Configuration is loaded from the environment. An optional `.env` file can be loaded with `WithEnvFile`.
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
## Usage
### STK Push (Lipa na M-Pesa Online)
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
    daraja.WithHTTPClient(customClient),
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
```
## Development
```bash
go test ./...
go vet ./...
```
CI runs on every push and pull request to `master` and `develop` via GitHub Actions.
## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
