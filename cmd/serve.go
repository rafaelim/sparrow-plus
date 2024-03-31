package cmd

import (
	"fmt"
	"log"
	"net/http"
	"sparrow-plus/api"
	"sparrow-plus/config"
)

func Serve() {
	router := http.NewServeMux()
	config := config.ReadConfig()

	SetupEnv(config)
	apiServe := api.NewAPIServe(fmt.Sprintf(":%v", config.Port))
	apiServe.Setup(router)

	log.Printf("Listening on port %v", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), router))
}
