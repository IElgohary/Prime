package main

import (
	"log"
	"net/http"

	"./bundles/person"
	"github.com/gorilla/mux"
)

func main() {
	// Controllers Declaration
	pc := &person.PersonController{}
	pc.Seed()
	router := mux.NewRouter()
	s := router.PathPrefix("/api/v1/").Subrouter()
	s.HandleFunc("/people", pc.GetPeople).Methods("GET")
	// s.HandleFunc("/people/{id}", getPerson).Methods("GET")
	// s.HandleFunc("/people/{id}", createPerson).Methods("POST")
	// s.HandleFunc("/people/{id}", deletePerson).Methods("DELETE")
	addr := ":1234"
    // http.HandleFunc("/", MyHandler)
    log.Println("listen on", addr)
    log.Fatal( http.ListenAndServe(addr, router) )
}
