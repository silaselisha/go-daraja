# Customer to Business (C2B) - Register URLs

Register validation and confirmation URLs for your shortcode.

```go
res, err := client.CustomerToBusiness(
  "https://example.com/confirm",
  "https://example.com/validate",
  daraja.COMPLETED, // or daraja.CANCELLED
)
```

- Uses your `DARAJA_BUSINESS_SHORT_CODE`.
- Endpoint: `mpesa/c2b/v1/registerurl`.
