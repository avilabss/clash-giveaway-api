package main

import (
	"fmt"
	"log"

	"github.com/git-avilabs/clash-giveaway-api/api"
	"github.com/git-avilabs/clash-giveaway-api/utils"
)

func main() {
	env, err := utils.LoadEnv(".")

	if err != nil {
		log.Fatal(fmt.Errorf("%s: %s", api.ErrFailedToLoadEnv.Error(), err))
	}

	server, err := api.NewServer(env)

	if err != nil {
		log.Fatal(fmt.Errorf("%s: %s", api.ErrFailedToSetupServer.Error(), err))
	}

	server.Router.Run(env.RunUri)
}
