package stripe_parser

import "net/http"

type Webhook struct {
}

func (receiver Webhook) handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
