package main

import (
	"net/http"
)



func handlerRediness(w http.ResponseWriter,r * http.Request) {
	respondWithJSON(w,200,struct{}{})
}




