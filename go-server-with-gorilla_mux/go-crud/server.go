package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Person struct {
	ID        string  `json:"id"`
	Firstname string  `json:"firstname"`
	Lastname  string  `json:"lastname"`
	Personal  *Detail `json:"detail"`
}

type Detail struct {
	Age    string `json:"age"`
	Number string `json:"number"`
}

var persons []Person

func getPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:idx], persons[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(persons)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Println("error Post", err)
	}
	person.ID = strconv.Itoa(rand.Intn(1000))
	persons = append(persons, person)
	json.NewEncoder(w).Encode(person)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range persons {
		if item.ID == params["id"] {
			persons = append(persons[:idx], persons[idx+1:]...)
			var person Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Println("error in put", err)
			}
			person.ID = params["id"]
			persons = append(persons, person)
			json.NewEncoder(w).Encode(person)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	persons = append(persons, Person{ID: "1", Firstname: "Dhana", Lastname: "22", Personal: &Detail{Age: "22", Number: "9090390390"}})

	r.HandleFunc("/persons", getPersons).Methods("GET")
	r.HandleFunc("/persons", createPerson).Methods("POST")
	r.HandleFunc("/persons/{id}", updatePerson).Methods("PUT")
	r.HandleFunc("/persons/{id}", deletePerson).Methods("DELETE")

	fmt.Printf("Strating server at port:8080")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
