package main

import "fmt"

var n int

//função init nada mais é do que uma função que será executada antes do main
//serve para configurações inicias e inicialização de variáveis
func init() {
	fmt.Println("Executou antes")
	n = 10
}

func main() {
	fmt.Println("Agora o main")
	fmt.Println(n)
}
