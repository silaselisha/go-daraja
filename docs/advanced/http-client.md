# Custom HTTP client

Inject a custom `*http.Client` to control timeouts, retries, proxies, and tracing.

```go
httpClient := &http.Client{ Timeout: 30 * time.Second }
client, err := daraja.NewClient(
  daraja.WithHTTPClient(httpClient),
)
```
