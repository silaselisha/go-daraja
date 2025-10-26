# Functional options

The client supports functional options to customize construction.

```go
client, err := daraja.NewClient(
  daraja.WithEnvFile("./example"),    // Load .env from path (optional)
  daraja.WithHTTPClient(httpClient),   // Inject custom *http.Client
  daraja.WithLogger(logger),           // Provide a logger with Debugf/Infof/Errorf
)
```

## Logger interface

```go
type Logger interface {
  Debugf(format string, args ...any)
  Infof(format string, args ...any)
  Errorf(format string, args ...any)
}
```
