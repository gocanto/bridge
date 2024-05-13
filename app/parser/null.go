package parser

import (
	"io"
	"net/http"
	"os"
)

type Null struct {
	request *http.Request
}

func (receiver Null) Handle(r *http.Request) error {
	return nil
}

func (receiver Null) PrintToConsole() {
	io.Copy(os.Stdout, receiver.request.Body)
}

func (receiver Null) GetRequestBody() io.ReadCloser {
	return receiver.request.Body
}
