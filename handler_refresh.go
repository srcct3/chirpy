package main

import (
	"net/http"
	"time"

	"github.com/srcct3/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerTokenRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}
	authToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid auth token", err)
		return
	}

	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), authToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token not found", err)
		return
	}

	token, err := auth.MakeJWT(user.ID, cfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Failed to generate jwt token", err)
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		Token: token,
	})
}

func (cfg *apiConfig) handlerTokenRevoke(w http.ResponseWriter, r *http.Request) {
	authToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "invalid auth token", err)
		return
	}
	_, err = cfg.db.RevokeRefreshToken(r.Context(), authToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't revoke refresh token", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
