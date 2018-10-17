package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person holds the information of a person
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address holds the address of the person
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

// People is a collection of persons
var people []Person

// GetPeople returns the list of all the persons present
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson returns the info of a single persons
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Printf("%+v \n", params)
	for _, item := range people {
		if item.ID == params["ID"] {
			if item.ID == params["ID"] {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// CreatePerson creates a new entry of person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Printf("Error decoding the person")
		return
	}
	person.ID = params["ID"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson deletes a new entry of person
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Printf("%+v", params)
	for index, item := range people {
		if params["ID"] == item.ID {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(r)
}

// Driver function
func main() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{ID}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{ID}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{ID}", DeletePerson).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8080", router))
}
