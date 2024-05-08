package bridge

import (
	"github.com/gocanto/bridge/app/entity"
	"net/http"
)

type Webhook struct {
	App entity.App
}

func Store(w http.ResponseWriter, r *http.Request) {

}
