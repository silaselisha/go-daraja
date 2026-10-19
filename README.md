<div align="center">
  <img src="public/images/gopokeamascott.png" alt="pokea" width="80" />

  # pokea

  Go client for Safaricom M-Pesa APIs.

  [![Go](https://img.shields.io/badge/go-%3E%3D1.22-00ADD8?logo=go&logoColor=white)](https://go.dev/)
  [![License](https://img.shields.io/github/license/silaselisha/pokea)](LICENSE)
</div>

## Error handling

Every method returns a Go `error` when a request fails: network errors, non-2xx responses, M-Pesa API `errorCode` payloads and non-JSON bodies. Errors are `*handler.APIError` values, so you can inspect them with `errors.As`. Use `errors.Is(err, handler.ErrUnreachable)` to detect cases where the M-Pesa API could not be reached or returned an unreadable body.

```go
res, err := client.NIPush("payment", "0712345678", 1)
if err != nil {
    var apiErr *handler.APIError
    if errors.As(err, &apiErr) {
        log.Printf("m-pesa error %s (status %d): %s", apiErr.Code, apiErr.Status, apiErr.Message)
    }
    return err
}
```

For backward compatibility, a failed call still returns the response with `ErrorCode` and `ErrorMessage` set. `res.Failed()` and `res.AsError()` are available when you only have the response.

## Migrating from Daraja names

The `Daraja*` names are deprecated and will be removed in the next major release. They still work for now:

| Deprecated | Use instead |
| --- | --- |
| `handler.NewDarajaClient` | `handler.NewPokeaClient` or `handler.NewClient` |
| `handler.Daraja` | `handler.Pokea` |
| `handler.DarajaClient` | `handler.PokeaClient` |
| `handler.DarajaResParams` | `handler.PokeaResParams` |
| `handler.DarajaErrorParams` | `handler.PokeaErrorParams` |
| `handler.DarajaAuth` | `handler.PokeaAuth` |
| `DARAJA_*` environment variables | `POKEA_*` (for example `DARAJA_CONSUMER_KEY` → `POKEA_CONSUMER_KEY`) |

When a `POKEA_*` variable is not set, its `DARAJA_*` counterpart is used, whether it comes from the environment or a `.env` file. The embedded response field is now `PokeaErrorParams`, so code that reads `res.DarajaErrorParams` directly must switch to `res.PokeaErrorParams`. Promoted fields such as `res.ErrorCode` are unaffected.
