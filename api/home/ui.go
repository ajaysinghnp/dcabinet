package ui

import "net/http"

// Home is the handler for the root endpoint
func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to DCabinet!"))
	}
}
