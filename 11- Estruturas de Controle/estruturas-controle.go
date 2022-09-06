package main

import "fmt"

func main() {

	//IF
	num := 10

	if num > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//IF INIT
	if outroNumero := num; outroNumero > 0 {
		fmt.Println("O número é maior que zero")
	} else {
		fmt.Println("O número é menor ou igual a zero")
	}

	//quando se cria uma variável no if/init, você limita ela para ser usada somente dentro do escopo
	//ela deixa de existir depois que o IF é executado
}
