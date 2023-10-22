package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/api-go-file/domain"
	"practice/api-go-file/domain/person"
)

func main() {
	personService, err := person.NewService("person.json")
	if err != nil {
		fmt.Println("Error trying to create person service")
		return
	}

	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("Error trying to decode body. Body shold be a json. Error: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusBadRequest)
				return
			}

			if person.ID <= 0 {
				http.Error(w, "Error trying to create person. ID shoud be a positive integer", http.StatusBadRequest)
			}

			// criar pessoa
			err = personService.Create(person)
			if err != nil {
				fmt.Printf("Error trying crete person: %s", err.Error())
				http.Error(w, "Error trying to create person", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			return
		}

		http.Error(w, "Not implemented", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", nil)

}
