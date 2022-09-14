package main

import (
	"database/sql"
	"fmt"
	"log"

	// usar pacote de forma implícita, coloca o underscore na frente do pacote
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConexao := "golang:golang@/cursoGolang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)
}
