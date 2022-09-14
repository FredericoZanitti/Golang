package banco

import (
	"database/sql"

	_ "github.com/gorilla/mux" // Driver de conexão com o MySQL
)

// Abre a conexão com o bando de dados
func Conectar(*sql.DB, error) {
	stringConexao := "golang:golang@/cursoGolang?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro := db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
