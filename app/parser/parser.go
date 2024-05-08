package parser

import "net/http"

type Parser interface {
	handle(w http.ResponseWriter, r *http.Request) error
}
