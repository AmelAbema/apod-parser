package main

import (
	"github.com/AmelAbema/apod-parser/internal/auth"
	"github.com/AmelAbema/apod-parser/internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.ApodDatum)

func middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

	}
}
