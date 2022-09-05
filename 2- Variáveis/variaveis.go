package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "kakakakakaka"
		variavel4 string = "babbabababab"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "teste de inferencia 1", "teste de inferencia 2"

	fmt.Println(variavel5)
	fmt.Println(variavel6)

	const constante1 string = "teste de constante"

	fmt.Println(constante1)
}
