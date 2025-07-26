package main

import (
	"log"
	"net/http"
	"time"
)

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting Broker Service on port: %s\n", webPort)

	srv := &http.Server{
		Addr:         ":" + webPort,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
}
