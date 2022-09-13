package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregar página de usuario"))
}

func main() {
	// HTTP é um protocolo de comunicação - base da comunicação web
	// modelo Cliente -> Servidor

	// rotas
	// URI identificador do recurso
	// Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuario", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
