# Troubleshooting

- Service unreachable (ErrorCode `500.003.1001`): The SDK maps network or non-JSON errors to a generic Daraja error with this code and message: "Service is currently unreachable. Please try again later." Check network connectivity, credentials, and Daraja service status.
- Invalid phone number: Numbers must match local `07` or `01` format before normalization.
- Auth fails: Ensure `DARAJA_CONSUMER_KEY` and `DARAJA_CONSUMER_SECRET` are set and `MPESA_ENVIRONMENT` is valid (`sandbox`/`production`).
