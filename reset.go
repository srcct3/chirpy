package main

import (
	"fmt"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerMetricsReset(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/palin; charset=utf-8")
	if cfg.paltform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		msg := "Reset is only allowed in dev environment"
		log.Print(msg)
		w.Write([]byte(msg))
		return
	}
	cfg.fileserveHit.Store(0)
	err := cfg.db.DeleteUsers(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed to reset db: %s", err)
		fmt.Fprintf(w, "Failed to reset db: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hits reset to 0 and db reset to initial state"))
}
