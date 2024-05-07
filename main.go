package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gocanto/bridge/app"
	"github.com/gocanto/bridge/app/entity"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("./config/app")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(fmt.Errorf("error reading config file:, %s", err))
		return
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	viper.WatchConfig()

	var app entity.App
	if err := viper.UnmarshalKey("app", &app); err != nil {
		fmt.Println(fmt.Errorf("failed to unmarshal into config: %v", err))
		return
	}

	var services []entity.Services
	if err := viper.UnmarshalKey("services", &services); err != nil {
		fmt.Println(fmt.Errorf("failed to unmarshal into config: %v", err))
		return
	}

	app.Services = &services

	fmt.Println(app, *app.Services)
	fmt.Println("---------")
}

func main() {

	//fmt.Println("--->", viper.Get("app.name"))
	// Config file found and successfully parsed

	http.HandleFunc("/", bridge.Handler)

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
