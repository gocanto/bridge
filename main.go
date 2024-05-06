package main

import (
	"errors"
	"fmt"
	"github.com/gocanto/bridge/app"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		var errorType viper.ConfigFileNotFoundError
		if errors.As(err, &errorType) {
			fmt.Println("config not found", errorType, err)
			return
		} else {
			fmt.Println("Config file was found but another error was produced")
			return
		}
	}

	fmt.Println("--->", viper.Get("app.name"))
	// Config file found and successfully parsed

	http.HandleFunc("/", bridge.Handler)

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
