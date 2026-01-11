package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) middlewareFileserveMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserveHit.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/palin; charset=utf-8")
	fmt.Fprintf(w, "file server hit: %d", cfg.fileserveHit.Load())
}
