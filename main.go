package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/srcct3/chirpy/internal/database"
)

type apiConfig struct {
	fileserveHit atomic.Int32
	db           *database.Queries
	paltform     string
}

func main() {
	godotenv.Load()
	mux := http.NewServeMux()
	port := "8080"

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Printf("Error failed to open database: %s", err)
		return
	}

	apiCfg := apiConfig{
		fileserveHit: atomic.Int32{},
		db:           database.New(db),
		paltform:     platform,
	}

	fs := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	fsHandler := apiCfg.middlewareFileserveMetrics(fs)
	mux.Handle("/app/", fsHandler)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerMetricsReset)

	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("POST /api/chirps", apiCfg.handlerChirpsCreate)
	mux.HandleFunc("GET /api/chirps", apiCfg.handlerChirpsGet)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("serving started at http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
