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
	fmt.Println("Starting server on", config.HTTPServer.Host+":"+string(config.HTTPServer.Port))
	server := &http.Server{
		Addr:    config.HTTPServer.Host + ":" + string(config.HTTPServer.Port),
		Handler: router,
	}
	// start server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server: ", err)
	}

}
