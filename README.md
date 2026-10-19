# Pokea

A Go client for the Safaricom M-Pesa APIs.

[![CI](https://github.com/silaselisha/pokea/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/silaselisha/pokea/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/silaselisha/pokea.svg)](https://pkg.go.dev/github.com/silaselisha/pokea)
[![Go Version](https://img.shields.io/badge/go-%3E%3D1.22-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/github/license/silaselisha/pokea)](LICENSE)

## Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Requirements](#requirements)
4. [Installation](#installation)
5. [Quick start](#quick-start)
6. [Configuration](#configuration)
7. [Usage](#usage)
8. [Error handling](#error-handling)
9. [Migrating from Daraja names](#migrating-from-daraja-names)
10. [Development](#development)
11. [Contributing](#contributing)
12. [License](#license)

## Overview

Pokea handles OAuth token acquisition, request signing, security credential generation and response decoding for the M-Pesa APIs, so you can start payments with a single method call. It targets both the sandbox and production environments.

## Features

| API | Method | Endpoint |
| --- | --- | --- |
| M-Pesa Express (STK push) | `NIPush` | `mpesa/stkpush/v1/processrequest` |
| Business to Customer (B2C) | `BusinessToConsumer` | `mpesa/b2c/v1/paymentrequest` |
| Customer to Business (C2B) URL registration | `CustomerToBusiness` | `mpesa/c2b/v1/registerurl` |
| Business to Business (B2B) Buy Goods | `BusinessBuyGoods` | `mpesa/b2b/v1/paymentrequest` |
| Business Express Checkout (USSD push) | `BusinessExpressCheckout` | `v1/ussdpush/get-msisdn` |

- Every method has a `…Ctx` variant that accepts a `context.Context` for cancellation and deadlines.
- Configuration comes from environment variables or an optional `.env` file.
- Failures are returned as typed Go errors (`*handler.APIError`).
- Sandbox and production certificates are embedded for generating security credentials.

## Requirements

- Go 1.22 or later
- A Safaricom developer account and an app with a consumer key and consumer secret

## Installation

```sh
go get github.com/silaselisha/pokea
```

## Quick start

Set your credentials in the environment or in a `.env` file (see [Configuration](#configuration)), then send an STK push:

```go
package main

import (
    "log"

    "github.com/silaselisha/pokea/pkg/handler"
)

func main() {
    client, err := handler.NewClient(handler.WithEnvFile("."))
    if err != nil {
        log.Fatalf("create client: %v", err)
    }

    res, err := client.NIPush("Order #1024", "0712345678", 1)
    if err != nil {
        log.Fatalf("stk push: %v", err)
    }

    log.Printf("accepted: %s", res.CustomerMessage)
}
```

`NewClient` loads the configuration and fetches an access token. It returns an error if either step fails.

## Configuration

### Environment variables

Settings are read from the OS environment. When a directory is passed with `WithEnvFile`, a `.env` file in that directory is also read if it exists.

| Variable | Required for | Description |
| --- | --- | --- |
| `MPESA_ENVIRONMENT` | All | `sandbox` (default) or `production` |
| `POKEA_CONSUMER_KEY` | All | App consumer key |
| `POKEA_CONSUMER_SECRET` | All | App consumer secret |
| `POKEA_BUSINESS_SHORT_CODE` | STK push, C2B, B2B | Paybill or till number |
| `POKEA_PASS_KEY` | STK push | Lipa na M-Pesa Online passkey |
| `POKEA_TRANSACTION_TYPE` | STK push | Defaults to `CustomerPayBillOnline` |
| `POKEA_CALL_BACK_URL` | STK push | URL that receives the STK push result |
| `POKEA_ACCOUNT_REF` | STK push | Account reference shown to the customer |
| `POKEA_INITIATOR_NAME` | B2C | API operator username |
| `POKEA_INITIATOR_PASSWORD` | B2C, B2B | API operator password, used to generate the security credential |
| `POKEA_BUSINESS_CONSUMR_PARTY_A` | B2C | Short code that sends B2C payments |
| `POKEA_BUSINESS_EXPRESS_CHECKOUT_SHORT_CODE` | Business Express Checkout | Primary short code for USSD push |

Example `.env`:

```sh
MPESA_ENVIRONMENT=sandbox
POKEA_CONSUMER_KEY=your-consumer-key
POKEA_CONSUMER_SECRET=your-consumer-secret
POKEA_BUSINESS_SHORT_CODE=174379
POKEA_PASS_KEY=your-pass-key
POKEA_CALL_BACK_URL=https://example.com/mpesa/callback
POKEA_ACCOUNT_REF=CompanyXLTD
```

### Client options

| Option | Description |
| --- | --- |
| `WithEnvFile(dir)` | Read a `.env` file from `dir` in addition to the OS environment |
| `WithHTTPClient(hc)` | Use a custom `*http.Client` for all requests. The default has a 15 second timeout |

```go
client, err := handler.NewClient(
    handler.WithEnvFile("."),
    handler.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
)
```

## Usage

Phone numbers must be in the local format `07XXXXXXXX` or `01XXXXXXXX`. They are converted to `2547XXXXXXXX` / `2541XXXXXXXX` before being sent.

### M-Pesa Express (STK push)

Prompts the customer to authorize a payment on their phone.

```go
res, err := client.NIPush("Order #1024", "0712345678", 100)
```

### Business to Customer (B2C)

Sends money from a business short code to a customer.

```go
res, err := client.BusinessToConsumer(
    500,                        // amount
    handler.BusinessPayment,    // SalaryPayment, BusinessPayment or PromotionalPayment
    "0712345678",               // customer phone number
    "Refund for order #1024",   // remarks
    "https://example.com/b2c/timeout",
    "https://example.com/b2c/result",
)
```

### Customer to Business (C2B) URL registration

Registers the URLs that receive C2B validation and confirmation requests. The response type controls what M-Pesa does when the validation URL is unreachable.

```go
res, err := client.CustomerToBusiness(
    "https://example.com/c2b/confirmation",
    "https://example.com/c2b/validation",
    handler.COMPLETED, // or handler.CANCELLED
)
```

### Business to Business (B2B) Buy Goods

Pays a till number from your business short code.

```go
res, err := client.BusinessBuyGoods(
    1000,                         // amount
    "apiuser",                    // initiator username
    "000000",                     // receiving short code
    "BusinessBuyGoods",           // command ID
    "Supplier payment",           // remarks
    "https://example.com/b2b/result",
    "https://example.com/b2b/timeout",
    "4",                          // receiver identifier type
    "4",                          // sender identifier type
    "INV-2048",                   // account reference
)
```

### Business Express Checkout (USSD push)

Sends a USSD push to a merchant to pay another business.

```go
res, err := client.BusinessExpressCheckout(
    "INV-2048",                       // payment reference
    "https://example.com/bec/callback",
    "Example Supplies",               // partner name
    "000000",                         // receiver short code
    1000,                             // amount
)
```

### Context support

`NewClient` returns a `*handler.PokeaClient`, which exposes a context-aware variant of every method:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

res, err := client.NIPushCtx(ctx, "Order #1024", "0712345678", 100)
```

`NewPokeaClient(dir)` is a shorthand for `NewClient(WithEnvFile(dir))` that returns the `handler.Pokea` interface. The interface does not include the `…Ctx` methods.

## Error handling

Every method returns a Go `error` when a request fails. That covers network errors, non-2xx responses, M-Pesa API `errorCode` payloads and non-JSON bodies. Errors are `*handler.APIError` values, so you can inspect them with `errors.As`. Use `errors.Is(err, handler.ErrUnreachable)` to detect cases where the M-Pesa API could not be reached or returned an unreadable body.

```go
res, err := client.NIPush("Order #1024", "0712345678", 100)
if err != nil {
    var apiErr *handler.APIError
    if errors.As(err, &apiErr) {
        log.Printf("m-pesa error %s (status %d): %s", apiErr.Code, apiErr.Status, apiErr.Message)
    }
    return err
}
```

`APIError` fields:

| Field | Description |
| --- | --- |
| `Status` | HTTP status code, or `0` if no response was received |
| `Code` | M-Pesa error code, for example `400.002.02` |
| `Message` | Error message from M-Pesa, or a description of the failure |
| `RequestID` | M-Pesa request ID, when provided |
| `RawBody` | Raw response body, when available |
| `Op` | The operation that failed, for example `pokea auth` or `pokea request` |

For backward compatibility, a failed call still returns the response with `ErrorCode` and `ErrorMessage` set. `res.Failed()` and `res.AsError()` are available when you only have the response.

## Migrating from Daraja names

The `Daraja*` names are deprecated and will be removed in the next major release. They still work for now:

| Deprecated | Use instead |
| --- | --- |
| `handler.NewDarajaClient` | `handler.NewPokeaClient` or `handler.NewClient` |
| `handler.Daraja` | `handler.Pokea` |
| `handler.DarajaClient` | `handler.PokeaClient` |
| `handler.DarajaResParams` | `handler.PokeaResParams` |
| `handler.DarajaErrorParams` | `handler.PokeaErrorParams` |
| `handler.DarajaAuth` | `handler.PokeaAuth` |
| `DARAJA_*` environment variables | `POKEA_*` (for example `DARAJA_CONSUMER_KEY` → `POKEA_CONSUMER_KEY`) |

When a `POKEA_*` variable is not set, its `DARAJA_*` counterpart is used, whether it comes from the environment or a `.env` file. The embedded response field is now `PokeaErrorParams`, so code that reads `res.DarajaErrorParams` directly must switch to `res.PokeaErrorParams`. Promoted fields such as `res.ErrorCode` are unaffected.

## Development

```sh
git clone https://github.com/silaselisha/pokea.git
cd pokea
go vet ./...
go test ./...
```

Some tests call the M-Pesa sandbox with the credentials in `example/.env`. They are skipped when the sandbox returns an error, so the suite passes offline. The remaining tests use a stubbed HTTP transport.

## Contributing

Contributions are welcome.

1. Open an issue to discuss bugs or larger changes, using the [issue templates](https://github.com/silaselisha/pokea/issues/new/choose).
2. Fork the repository and create a branch from `master`.
3. Add tests for your change and make sure `go vet ./...` and `go test ./...` pass.
4. Open a pull request describing the change.

Notable changes are recorded in the [changelog](CHANGELOG.md).

## License

Pokea is released under the [MIT License](LICENSE).
