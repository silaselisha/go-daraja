# Quickstart

The client reads configuration from environment variables (12‑factor). You may also provide a path to a directory containing a `.env` file using `WithEnvFile` or the backward‑compatible `NewDarajaClient`.

## Minimal example

```go
package main

import (
  "log"
  daraja "github.com/silaselisha/go-daraja/pkg/handler"
)

func main() {
  client, err := daraja.NewClient()
  if err != nil { log.Fatal(err) }

  // STK Push (NI)
  _, err = client.NIPush("payment for order #123", "0712345678", 100)
  if err != nil { log.Fatal(err) }
}
```

## With functional options

```go
httpClient := &http.Client{ Timeout: 20 * time.Second }
client, err := daraja.NewClient(
  daraja.WithHTTPClient(httpClient),
  daraja.WithEnvFile("./example"), // loads ./example/.env if present
)
```

## Backward‑compatible constructor

```go
client, err := daraja.NewDarajaClient("./example")
```
