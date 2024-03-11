package main

import (
	"fmt"
	"log"

	"github.com/silaselisha/go-daraja/util"
	"github.com/silaselisha/go-daraja/pkg/auth"
)

func main() {
	configs, err := util.LoadConfigs("./..")
	if err != nil {
		log.Fatal(err)
		return
	}
	consumerKey := configs.DarajaConsumerKey
	consumerSecret := configs.DarajaConsumerSecret
	authToken := util.GenAuthorizationToken(consumerKey, consumerSecret)
	fmt.Print(authToken)
	res, _ := auth.NewDarajaAuth(auth.URL, authToken)
	fmt.Println(res)
}
