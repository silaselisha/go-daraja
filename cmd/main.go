package main

import (
	"fmt"
	"log"

	"github.com/silaselisha/go-daraja/pkg/handler"
	"github.com/silaselisha/go-daraja/util"
)

func main() {
	configs, err := util.LoadConfigs("./..")
	if err != nil {
		log.Fatal(err)
		return
	}

	authToken := util.GenAuthorizationToken(configs.DarajaConsumerKey, configs.DarajaConsumerSecret)
	res, err := handler.NewDarajaAuth(handler.URL, authToken)
	if err != nil {
		log.Print(err)
		return
	}

	stk, err := res.MpesaExpress("Payment of X", "0708374149", 1)
	if err != nil {
		log.Print(err)
		return
	}
	
	fmt.Print(string(stk))
}
