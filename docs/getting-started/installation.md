# Installation

```bash
go get github.com/silaselisha/go-daraja
```

Then import and construct a client:

```go
import (
  daraja "github.com/silaselisha/go-daraja/pkg/handler"
)

func main() {
  client, err := daraja.NewClient()
  if err != nil { panic(err) }
  _ = client
}
```

## Requirements

- Go >= 1.22
- Network access to Safaricom Daraja endpoints
