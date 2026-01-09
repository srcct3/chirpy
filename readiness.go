package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain charset=utf-8")
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
