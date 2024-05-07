package entity

type Services struct {
	Name       string            `mapstructure:"name"`
	Parser     string            `mapstructure:"parser"`
	RateLimit  int16             `mapstructure:"rate_limit"`
	Identifier ServiceIdentifier `mapstructure:"identifier"`
}

type ServiceIdentifier struct {
	Id      string `mapstructure:"id"`
	Default string `mapstructure:"default"`
}
