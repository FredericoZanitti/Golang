package main

import "fmt"

//é o que mais se assemelha a uma classe em outras linguagens de OO
type usuario struct {
	nome  string
	idade int8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Frederico Zanitti"
	u.idade = 45
	fmt.Println(u)

	//popular por inferencia de tipos
	usuario2 := usuario{"Fred", 45}
	fmt.Println(usuario2)

	//caso não tenha todos os dados do struct, basta informar o campo
	usuario3 := usuario{nome: "Fred Zanitti"}
	fmt.Println(usuario3)
}
