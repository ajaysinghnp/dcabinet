package api

import (
	"net/http"

	handlers "github.com/ajaysinghnp/dcabinet/api/home"
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

	router.HandleFunc("/", handlers.Home)

	return router
}
