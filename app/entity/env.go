package entity

const PRODUCTION = "production"

type Env struct {
	Mode       string
	LogLevel   string
	ServerPort string
}

func (receiver *Env) IsProduction() bool {
	return receiver.Mode == PRODUCTION
}
