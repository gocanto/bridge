package stripe

import "net/http"

//https://docs.stripe.com/webhooks

type Webhook struct {
}

func (receiver Webhook) handle(w http.ResponseWriter, r *http.Request) error {
	return nil
}
