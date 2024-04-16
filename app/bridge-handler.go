package bridge

import (
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Extract service identifier from the URL path
	path := strings.Trim(r.URL.Path, "/")
	var redirectURL string

	switch path {
	case "service1":
		redirectURL = "http://example.com/service1"
	case "service2":
		redirectURL = "http://example.com/service2"
	default:
		// If the URL does not specify, use the header
		sender := r.Header.Get("X-Sender-ID")
		switch sender {
		case "sender1":
			redirectURL = "http://example.com/service1"
		case "sender2":
			redirectURL = "http://example.com/service2"
		default:
			http.Error(w, "Unauthorized sender or unknown service", http.StatusUnauthorized)
			return
		}
	}

	// Redirect to the determined URL
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
