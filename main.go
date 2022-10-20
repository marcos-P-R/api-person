package main

import (
	"api-cadastro-pessoas/rotas"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	rota := mux.NewRouter().StrictSlash(true)

	rota.HandleFunc("/", rotas.GetAll).Methods("GET")
	rota.HandleFunc("/persons", rotas.Create).Methods("POST")
	var port = ":3000"
	fmt.Println("Server running in port: ", port)
	log.Fatal(http.ListenAndServe(port, rota))
}
