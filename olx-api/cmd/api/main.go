package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kishanghosh090/GO-MONOLITH/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)

	// set server configs
	srv := http.Server{
		Addr:         ":8090",
		Handler:      mux,
		ReadTimeout:  time.Minute * 2,
		WriteTimeout: time.Minute * 5,
		IdleTimeout:  time.Second * 60,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
