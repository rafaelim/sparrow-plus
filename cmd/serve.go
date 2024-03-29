package cmd

import (
	"fmt"
	"log"
	"net/http"
	"sparrow-plus/api"
	"sparrow-plus/config"
)


func Serve() {
	ValidateEnv()
	config := config.ReadConfig()
	api.Setup()

	log.Printf("Listening on port %v", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.Port), nil))
}