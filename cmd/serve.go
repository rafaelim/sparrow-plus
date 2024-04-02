package cmd

import (
	"fmt"
	"log"
	"net/http"
	"sparrow-plus/api"
	"sparrow-plus/config"

	"github.com/rs/cors"
)

func Serve() {
	router := http.NewServeMux()
	config := config.ReadConfig()

	SetupEnv(config)
	database := SetupDatabase()
	apiServe := api.NewAPIServe("", database)
	apiServe.Setup(router)

	handler := cors.AllowAll().Handler(router)
	serve := &http.Server{
		Addr: fmt.Sprintf(":%v", config.Port),
		Handler: handler,
	}
	
	log.Printf("Listening on port %v", config.Port)
	log.Fatal(serve.ListenAndServe())
}
