package main

import (
	"fmt"
	"net/http"

	"github.com/muqeeth26832/go-rssagg/internal/auth"
	"github.com/muqeeth26832/go-rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error : %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Coudnt get user : %v", err))
			return
		}
		// if we got the user 
		// then do what the handler wanted to do
		handler(w,r,user)
	}
}
