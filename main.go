package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	port := "8080"

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("serving started at http://localhost:%s\n", port)
	log.Fatal(srv.ListenAndServe())
}
