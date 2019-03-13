package service

import (
	"encoding/json"
	"io"
	"net/http"
	"rest-api/src/structs"
)

var people []structs.Person

func init() {
	people = append(people, structs.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &structs.Address{City: "City X", State: "State X"}})
	people = append(people, structs.Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &structs.Address{City: "City Z", State: "State Y"}})
	people = append(people, structs.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func GetPeople() []structs.Person {
	return people
}

func GetPerson(params map[string]string) structs.Person {
	for _, item := range people {
		if item.ID == params["id"] {
			return item
		}
	}
	return structs.Person{}
}

func CreatePerson(body io.ReadCloser, params map[string]string) structs.Person {
	var person structs.Person
	_ = json.NewDecoder(body).Decode(&person)
	person.ID = params["id"]


	people = append(people, person)

	return person
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {}
