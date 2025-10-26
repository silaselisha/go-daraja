# Configuration

The client is configured via environment variables. A `.env` file in a specified directory may be read if you pass `WithEnvFile(path)` (or use `NewDarajaClient(path)`). OS environment variables always take precedence.

## Required

- `MPESA_ENVIRONMENT` — `sandbox` or `production` (default: `sandbox`)
- `DARAJA_CONSUMER_KEY`
- `DARAJA_CONSUMER_SECRET`

## Common (STK Push / C2B)

- `DARAJA_BUSINESS_SHORT_CODE`
- `DARAJA_PASS_KEY`
- `DARAJA_TRANSACTION_TYPE` (default: `CustomerPayBillOnline`)
- `DARAJA_CALL_BACK_URL`
- `DARAJA_ACCOUNT_REF`

## B2C

- `DARAJA_INITIATOR_NAME`
- `DARAJA_INITIATOR_PASSWORD`
- `DARAJA_BUSINESS_CONSUMR_PARTY_A` (your B2C PartyA short code)

## Business Express

- `DARAJA_BUSINESS_EXPRESS_CHECKOUT_SHORT_CODE`

## Optional

- `DARAJA_PARTY_A`, `DARAJA_PARTY_B`
- `DARAJA_CREDIT_PARTY_IDENTIFIER`

## Loading order

- Defaults applied by the loader
- Values from `.env` in the provided path (if the file exists)
- OS environment variables (override everything)
