package main

import "fmt"

func funcao1() {
	fmt.Println("Executa a função 1")
}

func funcao2() {
	fmt.Println("Executa a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("A média do aluno será calculada")
	fmt.Println("Entrando no calculo para ver se o aluno será aprovado")
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else {
		return false
	}
}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoEstaAprovado(10.0, 8.0))
}

//defer -> server para adiar a execução de uma função
