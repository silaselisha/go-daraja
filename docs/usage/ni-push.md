# STK Push (NI)

Initiate an STK Push to a customer's handset.

```go
res, err := client.NIPush(
  "payment for order #123", // description
  "0712345678",             // phone number (07/01 format)
  100,                       // amount
)
```

- Uses `DARAJA_BUSINESS_SHORT_CODE`, `DARAJA_PASS_KEY`, `DARAJA_CALL_BACK_URL`, `DARAJA_ACCOUNT_REF`, and `DARAJA_TRANSACTION_TYPE`.
- Phone numbers are normalized using the internal formatter; must match `07`/`01` local format.
- Endpoint: `mpesa/stkpush/v1/processrequest` on the selected base URL.
