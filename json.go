package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter,code int,msg string){
	if code >499 {
		log.Println("Responding with 5XX error :",msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w,code,errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter,code int,payload interface{}){
	data , err   := json.Marshal(payload)
	if err!=nil {
		// internal server error
		log.Printf("Faile to marshal JSON response :%v",payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type","application/json")
	// ok
	w.WriteHeader(code)
	w.Write(data)
}

