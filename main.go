package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Marcelo", LastName: "Scaldaferro", Address: &Address{City: "Montevideo", State: "Montevideo"}})
	people = append(people, Person{ID: "2", FirstName: "Guzman", LastName: "Scaldaferro", Address: &Address{City: "Montevideo", State: "Montevideo"}})
	people = append(people, Person{ID: "3", FirstName: "Andrea", LastName: "Paulet", Address: &Address{City: "Sauce", State: "Canelones"}})
	people = append(people, Person{ID: "4", FirstName: "Ethan", LastName: "Hughes", Address: &Address{City: "Toledo", State: "Canelones"}})
	people = append(people, Person{ID: "5", FirstName: "Irma", LastName: "Ceja"})
	people = append(people, Person{ID: "6", FirstName: "Walter", LastName: "Scaldaferro"})

	// endpoints
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

	holamundo = "dasdada"

	log.Println(holamundo)
}
