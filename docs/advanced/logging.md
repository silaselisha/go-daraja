# Logging

Provide a logger implementing the minimal interface to get observability into requests and responses.

```go
client, err := daraja.NewClient(
  daraja.WithLogger(myLogger),
)
```

The logger is optional. If omitted, the client operates silently.
