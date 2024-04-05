package cmd

import (
	"fmt"
	"log"
	"sparrow-plus/cmd/api"
	"sparrow-plus/config"
)

func Serve() {
	config := config.ReadConfig()

	SetupEnv(config)
	database := SetupDatabase()
	apiServe := api.NewAPIServe(fmt.Sprintf(":%v", config.Port), database)

	if err := apiServe.Run(); err != nil {
		log.Fatal(err)
	}
}
