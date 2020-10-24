package main

import (
	"log"
	"net/http"

	"github.com/Mynor2397/API-JWT/authentication"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("public")))
	mux.HandleFunc("/logeo", authentication.IndexLogin)
	mux.HandleFunc("/login", authentication.Login)
	mux.HandleFunc("/validate", authentication.ValidateToken)

	log.Println("Listen and serve in http://localhost:6060")
	http.ListenAndServe(":6060", mux)
}
