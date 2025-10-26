# Business to Consumer (B2C)

Send funds from a business short code to a consumer MSISDN.

```go
res, err := client.BusinessToConsumer(
  100,                 // amount
  daraja.SalaryPayment, // txnType (SalaryPayment|BusinessPayment|PromotionalPayment)
  "0712345678",       // customer number
  "salary for June",  // remarks
  "https://example.com/timeout", // queue timeout URL
  "https://example.com/result",  // result URL
)
```

- Requires `DARAJA_INITIATOR_NAME`, `DARAJA_INITIATOR_PASSWORD` to derive SecurityCredential (embedded certs).
- Also uses `DARAJA_BUSINESS_CONSUMR_PARTY_A`.
- Endpoint: `mpesa/b2c/v1/paymentrequest`.
