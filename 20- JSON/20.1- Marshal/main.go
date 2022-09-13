package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

// mapear um sturct para JSON usa a crase e a chave do jason
// quando for convertido, usa as chaves para comparar com o jason
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Pitbull", 6}
	fmt.Println(c)

	cachorroEmJASON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJASON)
	// imprimir em formato de JSON
	fmt.Println(bytes.NewBuffer(cachorroEmJASON))

	// outra forma, usando o map
	c2 := map[string]string{"nome": "Toby", "raca": "SRD"}
	fmt.Println(c2)

	cachorroEmJASON2, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEmJASON2)
	// imprimir em formato de JSON
	fmt.Println(bytes.NewBuffer(cachorroEmJASON2))
}
