package bridge

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HandlerResponse struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL, r.Body, r.RequestURI, &r.Response, r.Context(), &r)

	//payload := Payload{
	//	request: r,
	//	writer:  &w,
	//	body:    r.Body,
	//	content: content,
	//}
	//
	//content, err := r.Body.Read(payload.content)
	//if err != nil {
	//	return err
	//}
	//
	//// Extract service identifier from the URL path
	//path := strings.Trim(r.URL.Path, "/")
	//var redirectURL string
	//
	////r.Header
	//
	//switch path {
	//case "service1":
	//	redirectURL = "http://example.com/service1"
	//case "service2":
	//	redirectURL = "http://example.com/service2"
	//default:
	//	// If the URL does not specify, use the header
	//	sender := r.Header.Get("X-Sender-ID")
	//	switch sender {
	//	case "sender1":
	//		redirectURL = "http://example.com/service1"
	//	case "sender2":
	//		redirectURL = "http://example.com/service2"
	//	default:
	//		http.Error(w, "Unauthorized sender or unknown service", http.StatusUnauthorized)
	//		return err
	//	}
	//}

	w.Header().Set("Content-Type", "application/json")

	handlerResponse := HandlerResponse{Message: "Ok"}
	response, err := json.Marshal(handlerResponse)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("%s", err)))
		return
	}

	_, _ = w.Write(response)

	//return json.NewEncoder(w).Encode(payload)

	// Redirect to the determined URL
	//http.Redirect(w, r, redirectURL, http.StatusFound)
}
