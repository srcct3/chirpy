package main

import "net/http"

func (cfg *apiConfig) handlerMetricsReset(w http.ResponseWriter, _ *http.Request) {
	cfg.hit.Store(0)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/palin; charset=utf-8")
	w.Write([]byte("file server metrics has been set to 0"))
}
