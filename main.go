package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	hit atomic.Int32
}

func main() {
	mux := http.NewServeMux()
	port := "8080"

	apiCfg := apiConfig{
		hit: atomic.Int32{},
	}

	fs := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	fsHandler := apiCfg.middlewareFileserveMetrics(fs)
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerMetricsReset)

	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("serving started at http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
