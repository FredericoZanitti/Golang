package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CRUD - RELAÇÃO COM MÉTODOS
// CREATE - POST
// READ - GET
// UPDATE - PUT
// DELETE - DELETE

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuarios).Methods(http.MethodGet)

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
