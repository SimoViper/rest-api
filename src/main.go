package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest-api/src/gcpdatastore"
	"rest-api/src/service"
	"rest-api/src/structs"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people", CreatePerson).Methods("POST")
	//router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(service.GetPerson(params))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	var all = gcpdatastore.GetAll()
	json.NewEncoder(w).Encode(all)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var body  structs.Person
	_ = json.NewDecoder(r.Body).Decode(&body)
	gcpdatastore.SavePerson(body)

	var person = service.CreatePerson(r.Body, params)
	json.NewEncoder(w).Encode(person)
}
