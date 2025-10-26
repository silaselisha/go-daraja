<div align="center" style="margin-bottom: 0px!important; padding: 0px;">
  <img src="./public/images/godarajamascott.png" alt="go-daraja logo" height="100px"/>
</div>

<div style="align-items: center; margin-top: 0px !important; margin-bottom: 14px;" align="center">
  <p style="text-align: center;" align="center">
    <img src="https://img.shields.io/badge/go-%3E%3D1.22-blue?logo=go" />
    <img src="https://img.shields.io/github/license/silaselisha/go-daraja" />
    <img src="https://img.shields.io/github/actions/workflow/status/silaselisha/go-daraja/ci.yml?label=CI" />
  </p>
</div>

<div style="text-align: center;">
  <h1>
    Introducing the Ultimate Go Daraja (M‑Pesa) Client
  </h1>
</div>

go-daraja is a lightweight, idiomatic Go client for Safaricom's Daraja (M‑Pesa) APIs. It emphasizes a clean developer experience, easy configuration across environments, and safe defaults for production.

### Why Would I use this?

- Easy to set up and install
- Have a clean, consistent API surface
- Context-aware calls for timeouts/cancellation
- Pluggable `http.Client` for custom transports/retries
- 12‑factor configuration via OS env, optional `.env`
- Embedded X.509 certs and typed error model
- Covers core Daraja flows (STK Push, B2C, C2B, B2B, Business Express)

## Table of Contents

- [Install](#install)
- [APIs Covered](#apis-covered)
- [Advanced Features](#advanced-features)
- [Configuration](#configuration)
- [Usage Examples](#usage-examples)
- [Project Structure](#project-structure)
- [Development](#development)
- [GitHub Stats](#github-stats)
- [Contributing](#contributing)
- [Security](#security)
- [License](#license)

<a id="install"></a>

<h2>
  Install
</h2>

### Go Get
```bash
go get github.com/silaselisha/go-daraja
```

Then import and construct a client:

```go
import (
  "log"
  daraja "github.com/silaselisha/go-daraja/pkg/handler"
)

func main() {
  client, err := daraja.NewClient()
  if err != nil { log.Fatal(err) }
  // ...
}
```

<a id="apis-covered"></a>

<h2>
  APIs Covered
</h2>

- STK Push (NI)
- Business to Customer (B2C)
- Customer to Business (C2B)
- Business to Business (B2B)
- Business Express Checkout
- OAuth access token generation

<a id="advanced-features"></a>

<h2>
  Advanced Features
</h2>

Use functional options with `NewClient(...)` to tailor behavior:

- HTMX-style context support via `*_Ctx` methods for cancellation/timeouts
- Inject your own `http.Client` (retries, proxies, tracing)
- Load configuration from `.env` with `WithEnvFile("path")`
- Embedded sandbox/production X.509 certs via `go:embed`
- Consistent, typed errors; backward compatible responses

<a id="configuration"></a>

<h2>
  Configuration
</h2>

Set required envs (12‑factor friendly). `.env` is optional and supported.

```bash
export MPESA_ENVIRONMENT=sandbox
export DARAJA_CONSUMER_KEY=your_key
export DARAJA_CONSUMER_SECRET=your_secret
export DARAJA_BUSINESS_SHORT_CODE=174379
export DARAJA_PASS_KEY=your_pass_key
export DARAJA_CALL_BACK_URL=https://your.app/callback
export DARAJA_ACCOUNT_REF=YourApp
# Required for B2C/B2B:
export DARAJA_INITIATOR_NAME=your_initiator
export DARAJA_INITIATOR_PASSWORD=your_password
```

<a id="usage-examples"></a>

<h2>
  Usage Examples
</h2>

Quickstart:

```go
client, err := daraja.NewClient()
if err != nil { log.Fatal(err) }
```

STK Push (NI):

```go
res, err := client.NIPush("payment for order #123", "0712345678", 100)
```

B2C:

```go
res, err := client.BusinessToConsumer(
  100, 0, "0712345678", "salary",
  "https://example.com/timeout",
  "https://example.com/result",
)
```

C2B register URLs:

```go
res, err := client.CustomerToBusiness(
  "https://example.com/confirm",
  "https://example.com/validate",
  1,
)
```

Business Express Checkout:

```go
res, err := client.BusinessExpressCheckout(
  "PAY-123",
  "https://example.com/callback",
  "YourBiz",
  "600000",
  100,
)
```

Backward compatible constructor:

```go
client, err := daraja.NewDarajaClient(".")
```

<a id="project-structure"></a>

<h2>
  Project Structure
</h2>

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

<a id="development"></a>

<h2>
  Development
</h2>

Prereqs: Go >= 1.22

- Run tests:
```bash
go test ./...
```

- Lint/inspect locally:
```bash
go vet ./...
```

<a id="github-stats"></a>

<h2>
  GitHub Stats
</h2>

<p align="center">
  <img src="https://repobeats.axiom.co/api/embed/36b264b4be024052073f9c5703b102cd24693c62.svg" alt="go-daraja stats" title="Repobeats analytics image"/>
 </p>

<a id="contributing"></a>

<h2>
  Contributing
</h2>

Contributions are welcome! Please read `CONTRIBUTING.md` and open an issue or PR.

<a id="security"></a>

<h2>
  Security
</h2>

If you discover a security vulnerability, please follow the process in `SECURITY.md`.

<a id="license"></a>

<h2>
  License
</h2>

This project is licensed under the MIT License. See `LICENSE` for details.
