package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/kishanghosh090/GO-MONOLITH/internal/config"
	"github.com/kishanghosh090/GO-MONOLITH/internal/handlers"
)

func main() {
	config.MustLoad()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.Health)

	// set server configs
	srv := http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      mux,
		ReadTimeout:  time.Minute * 2,
		WriteTimeout: time.Minute * 5,
		IdleTimeout:  time.Second * 60,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}

}
