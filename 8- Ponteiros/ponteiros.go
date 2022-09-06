package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	//a partir desse ponto as variaveis 1 e 2 não tem mais nenhuma ligação uma com a outra
	//se mudar uma a outra não será afetada

	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)
	variavel2 *= 3
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro, *ponteiro)
	//pegar o valor do ponteiro: *ponteiro

}
