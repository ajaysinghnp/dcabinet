package api

import (
	"net/http"

	navigation "github.com/ajaysinghnp/dcabinet/api/navigation"
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

	// root handler
	router.HandleFunc("/", navigation.Home())

	return router
}
