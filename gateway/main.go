package main

import (
	"log"
	"net/http"
)

var (
	httpAddr = common.Evn("HTTP_ADDR", ":3000")
)

// main is the entry point for the application.

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
