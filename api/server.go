package api

import (
	"net/http"
)

func NewServer(addr string, router *http.ServeMux) *http.Server {
	// setup server
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	return server
}

func NewRouter() *http.ServeMux {
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to DCabinet!"))
	})

	return router
}
