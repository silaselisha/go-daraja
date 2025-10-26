# Business Express Checkout

Initiate a USSD push to collect MSISDN and proceed with payment.

```go
res, err := client.BusinessExpressCheckout(
  "PAY-123",                     // paymentRef
  "https://example.com/callback",// callback URL
  "YourBiz",                     // partner name
  "600000",                      // receiver short code
  100,                            // amount
)
```

- Uses `DARAJA_BUSINESS_EXPRESS_CHECKOUT_SHORT_CODE` as the primary shortcode.
- Endpoint: `/v1/ussdpush/get-msisdn`.
