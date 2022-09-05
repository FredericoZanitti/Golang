package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	altura    int16
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"João", "Silva", 23, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "Adminstração", "Newton Paiva"}
	fmt.Println(e1)
	fmt.Println(e1.nome)

	e2 := estudante{pessoa{"Frederico", "Zanitti", 45, 171}, "Sistema da Informação", "Newton Paiva"}
	fmt.Println(e2)
	fmt.Println(e2.nome)
	fmt.Println(e2.sobrenome)
}
