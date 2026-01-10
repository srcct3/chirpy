package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Body string `json:"body"`
	}

	type response struct {
		CleanedBody string `json:"cleaned_body"`
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

	cleaned := getCleanedBody(param.Body)
	respondWithJSON(w, http.StatusOK, response{
		CleanedBody: cleaned,
	})
}

func getCleanedBody(msg string) string {
	wordList := strings.Split(msg, " ")
	disAllowed := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	for i, word := range wordList {
		_, isProfane := disAllowed[strings.ToLower(word)]
		if isProfane {
			wordList[i] = "****"
		}
	}
	return strings.Join(wordList, " ")
}
