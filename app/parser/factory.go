package parser

type Factory struct {
}

func MakeParser() (Factory, error) {
	return Factory{}, nil
}
