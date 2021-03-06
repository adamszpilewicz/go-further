package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version = "1.0.0"
)

type config struct {
	port int
	env string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "(development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthCheck)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d",app.config.port),
		Handler:      mux,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}