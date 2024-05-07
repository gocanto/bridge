package entity

type App struct {
	Name     string      `mapstructure:"name"`
	Version  string      `mapstructure:"version"`
	Services *[]Services `mapstructure:"services"`
}
