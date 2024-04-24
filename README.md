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

## About

Go-daraja is an open-source project facilitating seamless integration of Safaricom's Daraja API into Golang applications, providing developers with a straightforward interface. It simplifies API interaction, allowing developers to focus on core application logic while encouraging community collaboration for ongoing improvement.

## Install and use go-daraja

1. Ensure that go is installed. In a situation where go is not installed click **[install go](https://go.dev/doc/install)**
2. Create a new go project

    ``` cmd
        mkdir duka-letu
        cd duka-letu
        go mod init github.com/user-github-name/project-name
    ```

3. Install go-daraja package

    ```go
       go get github.com/silaselisha/go-daraja
    ```
    
---
**Note**:
Majority of the services provided by daraja API require the client to be authenticated before invoking them. With go-daraja a simplified interface `Daraja` shall provide the user with all necessary services such as `ClientAuth`, primarily used to generate an `access-token`. 

``` go
    import (
	    "log"
	    daraja "github.com/silaselisha/go-daraja/pkg/handler"
    )

    func main() {
        // create a client
        client, err := daraja.NewDarajaClient(".")
        if err != nil {
            log.Panic(err)
        }

        // generate an access token
        auth, err := client.ClientAuth()
        if err != nil {
            log.Panic(err)
        }

        // invoke STK/NI Push
        buff, err := client.NIPush("test STK push", "0708374149", 1, auth.AccessToken)
        if err != nil {
            log.Panic(err)
        }
        log.Print(string(buff))
    }
```