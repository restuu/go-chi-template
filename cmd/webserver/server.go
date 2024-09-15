package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func initServer(handler http.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", port),
		Handler:           handler,
		ReadHeaderTimeout: 3 * time.Second,
	}

	return srv
}
