package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gocanto/bridge/app/entity"
	"github.com/spf13/viper"
)

func makeEnv(env *entity.Env) error {
	loader := viper.New()
	loader.AddConfigPath(".")
	loader.SetConfigName(".env")
	loader.SetConfigType("env")

	if err := loader.ReadInConfig(); err != nil {
		return err
	}

	env.Mode = loader.GetString("APP_ENV")
	env.LogLevel = loader.GetString("APP_LOG_LEVEL")

	watchFile(loader, &environment)

	return nil
}

func makeAppConfig(env *entity.Env, app *entity.App) error {
	loader := viper.New()
	loader.AddConfigPath(".")
	loader.SetConfigName("./config/app")
	loader.SetConfigType("yaml")

	if err := loader.ReadInConfig(); err != nil {
		return err
	}

	watchFile(loader, env)

	if err := loader.UnmarshalKey("app", &app); err != nil {
		return fmt.Errorf("failed to unmarshal into config [app]: %v", err)
	}

	var services []entity.Services
	if err := loader.UnmarshalKey("services", &services); err != nil {
		return fmt.Errorf("failed to unmarshal into config [services]: %v", err)
	}

	app.Services = &services
	app.Environment = env

	return nil
}

func watchFile(loader *viper.Viper, env *entity.Env) {
	if env.IsProduction() {
		return
	}

	loader.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	loader.WatchConfig()
}
