package entity

type App struct {
	Name     string `mapstructure:"name"`
	Version  int32  `mapstructure:"version"`
	LogLevel string `mapstructure:"log_level"`
	// ----
	Environment *Env
	Services    *[]Services `mapstructure:"services"`
}
