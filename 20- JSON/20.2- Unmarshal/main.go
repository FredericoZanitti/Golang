package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJASON := `{"nome":"Rex","raca":"Pitbull","idade":6}`

	var c cachorro

	// é necessário passar o JSON como um slice de byte. Ele não aceita o formato "legível"
	if erro := json.Unmarshal([]byte(cachorroEmJASON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	carrocho2EmJASON := `{"nome":"Toby","raca":"SRD"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(carrocho2EmJASON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
