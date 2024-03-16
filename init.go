package main

import (
	"fmt"
	"log"

	"github.com/silaselisha/go-daraja/util"
)

func init() {
	envs, err := util.LoadConfigs(".")
	if err != nil {
		log.Fatal(err)
	}
	util.Envs = envs

	fmt.Println("loaded")
}
