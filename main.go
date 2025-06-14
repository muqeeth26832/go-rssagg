package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {

	env := EnvConfig()


	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))


	v1Router := chi.NewRouter()

	v1Router.Get("/healthz",handlerRediness)
	v1Router.Get("/err",handlerErr)

	router.Mount("/v1",v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + env.PORT,
	}


	fmt.Println("Starting Server at PORT:",env.PORT)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalf("server issue :%e", err)
	}

}
