package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/Mynor2397/virtual-parish-office/internal/lib"
	"github.com/Mynor2397/virtual-parish-office/internal/routers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = lib.Config().PORT
	}

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"})
	originsOk := handlers.AllowedOrigins([]string{"*"})

	fmt.Printf("\nListen and serve on port:%s...\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handlers.CORS(headersOk, methodsOk, originsOk)(routers.InitRoutes())); err != nil {
		log.Panic(err)
	}
}
