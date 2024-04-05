package cmd

import (
	"fmt"
	"log"
	"sparrow-plus/cmd/api"
	configs "sparrow-plus/config"
)

func Serve() {
	SetupEnv()
	database := SetupDatabase(configs.Envs.DBName)
	apiServe := api.NewAPIServe(fmt.Sprintf(":%v", configs.Envs.Port), database)

	if err := apiServe.Run(); err != nil {
		log.Fatal(err)
	}
}
