package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 2 {
		i++
		fmt.Println("Incrementando i: ", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 7; j += 3 {
		fmt.Println("Incrementando j: ", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Maria", "José"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//inibindo um dos parametros
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//se usar range em uma string, o retorno será os códigos da tabela ASCII
	for indice, nome := range "FREDERICO" {
		fmt.Println(indice, nome, string(nome))
	}

	//for em um map
	usuario := map[string]string{
		"nome":      "Frederico",
		"sobrenome": "Zanitti",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//não é possível fazer loop em STRUCT
}
