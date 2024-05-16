package parser

import (
	"errors"
	"fmt"
	"github.com/gocanto/bridge/app/parser/stripe"
	"net/http"
	"strings"
)

type Factory struct {
	request       *http.Request
	requestHeader http.Header
	parser        Webhook
}

func MakeFactory(request *http.Request) (Factory, error) {
	factory := Factory{}
	parser, err := factory.getParserFrom(request)

	if err != nil {
		return factory, err
	}

	factory.request = request
	factory.parser = parser
	factory.requestHeader = request.Header

	return factory, nil
}

func (receiver Factory) getParserFrom(request *http.Request) (Webhook, error) {
	switch {
	case receiver.isStripe():
		var parser stripe.Stripe

		if err := stripe.MakeParser(request, &parser); err != nil {
			return stripe.Stripe{}, nil
		}

		return parser, nil
	case receiver.isLocalHost():
		return Null{}, errors.New("working with localhost")
	default:
		return Null{}, errors.New("unsupported parser")
	}
}

func (receiver Factory) isStripe() bool {
	header := receiver.requestHeader

	headerName := header.Get(stripe.HeaderName)
	headerValue := header.Get(stripe.AgentSeed)

	return receiver.serviceIs(headerName, headerValue)
}

func (receiver Factory) isLocalHost() bool {
	return receiver.serviceIs("", "localhost")
}

func (receiver Factory) serviceIs(headerName string, HeaderValue any) bool {
	signature := strings.TrimSpace(
		headerName,
	)

	if len(signature) > 0 {
		return true
	}

	agent := strings.TrimSpace(
		strings.ToLower(receiver.requestHeader.Get("User-Agent")),
	)

	return strings.Contains(agent, fmt.Sprint(HeaderValue))
}

func (receiver Factory) Parser() Webhook {
	return receiver.parser
}
