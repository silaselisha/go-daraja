package main

import (
	"github.com/silaselisha/go-daraja/pkg/auth"
)

func main() {
	auth.NewDarajaAuth(auth.URL, auth.AUTHORIZATION_TOKEN)
}
