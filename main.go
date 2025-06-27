package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajaysinghnp/dcabinet/lib/config"
)

func main() {
	// load config
	config := config.MustLoad()

	// database setup

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to DCabinet!"))
	})

	// setup server
	http_addr := config.HTTPServer.Host + ":" + config.HTTPServer.Port

	server := &http.Server{
		Addr:    http_addr,
		Handler: router,
	}
	// start server
	fmt.Println("Starting server on", http_addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
