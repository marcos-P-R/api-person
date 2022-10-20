package rotas

import (
	"encoding/json"
	"io"
	"net/http"
)

type Person struct {
	Id   int
	Name string
}

var Persons = []Person{
	{Id: 1, Name: "Marcos"},
	{Id: 2, Name: "Ana"},
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Persons)
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var P Person

	Body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(Body, &P); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	json.Unmarshal(Body, &P)

	Persons = append(Persons, P)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(P); err != nil {
		panic(err)
	}
}
