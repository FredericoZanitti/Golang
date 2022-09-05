package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//no Go podemos ter funções com mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtrai := n1 - n2
	return soma, subtrai
}

func main() {
	soma := somar(6, 7)
	fmt.Println(soma)

	resultSoma, resultSubt := calculosMatematicos(10, 22)
	fmt.Println(resultSoma, resultSubt)

	//no Go não é possível declarar uma variavel e não usar
	//então se quiser chamar a função calculosMatematicos mas por exemplo
	//não quiser usar o resutlado da subtração, basta passar um underscore _ no lugar
	//do retorno indesejado

	resultSoma2, _ := calculosMatematicos(20, 33)
	fmt.Println(resultSoma2)
}
