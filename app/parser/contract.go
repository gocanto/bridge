package parser

import (
	"io"
	"net/http"
)

type Webhook interface {
	Handle(r *http.Request) error
	PrintToConsole()
	GetRequestBody() io.ReadCloser
}
