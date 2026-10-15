<div align="center">
  <img src="docs/assets/images/godarajamascott.png" alt="go-daraja" width="80" />

  # go-daraja

  Go client for Safaricom Daraja (M-Pesa) APIs.

  [![Go](https://img.shields.io/badge/go-%3E%3D1.22-00ADD8?logo=go&logoColor=white)](https://go.dev/)
  [![License](https://img.shields.io/github/license/silaselisha/go-daraja)](LICENSE)
</div>

## Install

```bash
go get github.com/silaselisha/go-daraja
```

```go
client, err := daraja.NewClient()
```

## Config

Set environment variables:

```bash
export MPESA_ENVIRONMENT=sandbox
export DARAJA_CONSUMER_KEY=your_key
export DARAJA_CONSUMER_SECRET=your_secret
export DARAJA_BUSINESS_SHORT_CODE=174379
export DARAJA_PASS_KEY=your_pass_key
export DARAJA_CALL_BACK_URL=https://your.app/callback
```

## Usage

```go
// STK Push
res, err := client.NIPush("payment", "0712345678", 100)

// B2C
res, err := client.BusinessToConsumer(100, 0, "0712345678", "salary", timeoutURL, resultURL)
```

Supports STK Push, B2C, C2B, B2B, and Business Express.

## Docs

[https://silaselisha.github.io/go-daraja/](https://silaselisha.github.io/go-daraja/)

## License

MIT
