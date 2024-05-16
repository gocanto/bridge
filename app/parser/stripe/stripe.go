package stripe

import (
	"io"
	"net/http"
	"os"
)

const AgentSeed = "stripe.com"
const HeaderName = "Stripe-Signature"

type Stripe struct {
	request *http.Request
}

func MakeParser(request *http.Request, target *Stripe) error {
	(*target).request = request

	return nil
}

func (receiver Stripe) Handle(r *http.Request) error {
	return nil
}

func (receiver Stripe) PrintToConsole() {
	io.Copy(os.Stdout, receiver.request.Body)
}

func (receiver Stripe) GetRequestBody() io.ReadCloser {
	return receiver.request.Body
}
