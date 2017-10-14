package main

import (
	"log"
	"net/http"

	"./bundles/env"
	"./bundles/wolfram_alpha"
	"github.com/gorilla/mux"
)

func main() {
	// Controllers Declaration

	wa := &wolfram_alpha.WolframAlphaController{}
	env.New()
	router := mux.NewRouter()
	s := router.PathPrefix("/api/v1/").Subrouter()

	s.HandleFunc("/query", wa.HandleQuery).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+env.GetString("SERVER_PORT"), router))
}
