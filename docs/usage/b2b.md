# Business to Business (B2B) - BuyGoods

Transfer funds between short codes.

```go
res, err := client.BusinessBuyGoods(
  100,                 // amount
  "initiator",        // username (Initiator)
  "600000",           // receiver short code
  "BusinessPayBill",  // command ID (per API spec)
  "invoice #123",     // remarks
  "https://example.com/result",
  "https://example.com/timeout",
  "4",                // receiver identifier type
  "4",                // sender identifier type
  "ACC-REF-123",      // account reference
)
```

- Generates SecurityCredential from embedded certs.
- Uses your `DARAJA_BUSINESS_SHORT_CODE` as PartyA.
- Endpoint: `mpesa/b2b/v1/paymentrequest`.
