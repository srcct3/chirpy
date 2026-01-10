package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Body string `json:"body"`
	}

	type response struct {
		Valid bool `json:"valid"`
	}

	validCharLen := 140

	param := parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&param)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to decode json", err)
		return
	}

	if len(param.Body) > validCharLen {
		respondWithError(w, http.StatusBadRequest, "Chirp too long", nil)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Valid: true,
	})
}
