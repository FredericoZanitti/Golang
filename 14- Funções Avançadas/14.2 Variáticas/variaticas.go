package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, valor := range numeros {
		fmt.Println(texto, valor)
	}
}

func main() {
	totalSoma := soma(1, 11)
	fmt.Println(totalSoma)

	escrever("Olá Mundo!", 5, 4, 1, 8, 10, 2)
}

//não pode haver mais de um parâmetro variático na função
//e o variático deve obrigatoriamente ser o último
