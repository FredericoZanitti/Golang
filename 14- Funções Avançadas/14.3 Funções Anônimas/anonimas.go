package main

import "fmt"

func main() {

	//funções anôinimas não tem nome, e é definida pelo () no final da chave
	//ela é executada assim que é criada
	func() {
		fmt.Println("Olá meu pequeno camaleão")
	}()

	//com parametros
	func(texto string) {
		fmt.Println(texto)
	}("Passando parâmetros")

	//com retorno
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 20)
	}("Passando parâmetros")

	fmt.Println(retorno)
}
