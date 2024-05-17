package main

import (
	"fmt"
	"github.com/gocanto/bridge/app"
	"github.com/gocanto/bridge/app/entity"
	"log"
	"net/http"
)

var environment entity.Env
var application entity.App

func init() {
	if err := makeEnv(&environment); err != nil {
		log.Fatal(fmt.Errorf("error reading ENV config file:, %s", err))
	}

	if err := makeAppConfig(&environment, &application); err != nil {
		log.Fatal(fmt.Errorf("error reading the config file:, %s", err))
	}

	fmt.Println("---- Init ----")
	fmt.Println("--> App:", application)
	fmt.Println("--> Services:", *application.Services)
	fmt.Println("--> Env:", *application.Environment)
	fmt.Println("---------")
}

func main() {
	http.HandleFunc("/", bridge.Store)

	log.Printf("Server starting on port %s", environment.ServerPort)

	if err := http.ListenAndServe(environment.ServerPort, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
