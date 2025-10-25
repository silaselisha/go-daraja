<div align="center" style="margin-bottom: 0px!important; padding: 0px;">
    <img src="./public/images/godarajamascott.png" alt="godaraja logo" height="100px"/>
</div>

<div style="align-items: center; margin-top: 0px !important; margin-bottom: 14px;" align="center">
    <p style="text-align: center;" align="center">
        <img src="https://img.shields.io/badge/logo-go-blue?logo=go">
        <img src="https://img.shields.io/badge/logo-circleci-black?logo=circleci">
        <img src="https://img.shields.io/badge/logo-git-orange?logo=git">
        <img src="https://img.shields.io/badge/logo-markdown-skyblue?logo=markdown">
    </p>
    <h1 style="font-size: 48px; font-weight: 800; padding: 0px;">Go-daraja</h1>
</div>

## Table of Contents

- [About](#about)
- [Install](#install)
- [Usage Example](#usage-example)
- [GitHub Stats](#stats)

## About

Go-daraja is an open-source project facilitating seamless integration of Safaricom's Daraja API into Golang applications, providing developers with a straightforward interface. It simplifies API interaction, allowing developers to focus on core application logic while encouraging community collaboration for ongoing improvement.

## Install

1. Ensure that go is installed. In a situation where go is not installed click **[install go](https://go.dev/doc/install)**
2. Create a new go project

    ``` cmd
        mkdir duka-letu
        cd duka-letu
        go mod init github.com/user-github-name/project-name
    ```

    ```cmd
       cp example/.env .env
    ```

3. Install go-daraja package

    ```go
       go get "github.com/silaselisha/go-daraja"
    ```

## Usage Example

### Quick start (env-based)

Set the required environment variables (12-factor friendly). `.env` is optional.

Required for auth: `DARAJA_CONSUMER_KEY`, `DARAJA_CONSUMER_SECRET`. Common values:

```bash
export MPESA_ENVIRONMENT=sandbox
export DARAJA_CONSUMER_KEY=your_key
export DARAJA_CONSUMER_SECRET=your_secret
export DARAJA_BUSINESS_SHORT_CODE=174379
export DARAJA_PASS_KEY=your_pass_key
export DARAJA_CALL_BACK_URL=https://your.app/callback
export DARAJA_ACCOUNT_REF=YourApp
```

```go
import (
    "log"
    daraja "github.com/silaselisha/go-daraja/pkg/handler"
)

func main() {
    // New option-based constructor; loads OS env and optional .env automatically
    client, err := daraja.NewClient()
    if err != nil {
        log.Panic(err)
    }

    res, err := client.NIPush("test STK push", "0708374149", 1)
    if err != nil {
        log.Panic(err)
    }
    log.Printf("%+v\n", res)
}
```

### Backward-compatible constructor

The original constructor remains available and now treats `.env` as optional:

```go
client, err := daraja.NewDarajaClient(".") // looks for .env in current dir, also reads OS env
```

### Advanced: provide custom HTTP client or config

```go
httpClient := &http.Client{Timeout: 20 * time.Second}
client, err := daraja.NewClient(
    daraja.WithHTTPClient(httpClient),
    // daraja.WithEnvFile("./configs"),
    // daraja.WithConfig(daraja.Configs{ /* explicit values */ }),
)
```

### Usage Example

A sample implementation of MPESA STK push integration in a Go project

```go
    import (
        "log"
        daraja "github.com/silaselisha/go-daraja/pkg/handler"
    )

    func main() {
        // create a client by passing the path to your .env file
        client, err := daraja.NewDarajaClient(".")
        if err != nil {
            log.Panic(err)
        }

        // invoke STK/NI Push
        res, err := client.NIPush("test STK push", "0708374149", 1)
        if err != nil {
            log.Panic(err)
        }
        log.Printf("%+v\n", res)
    }
```

**Note**:
Majority of the services provided by daraja API require the client to be authenticated before invoking them. With go-daraja a simplified interface `Daraja` shall provide the user with all necessary services.

## Stats

<p align="center">
    <img src="https://repobeats.axiom.co/api/embed/36b264b4be024052073f9c5703b102cd24693c62.svg" alt="go-daraja stats" title="Repobeats analytics image"/>
</p>
