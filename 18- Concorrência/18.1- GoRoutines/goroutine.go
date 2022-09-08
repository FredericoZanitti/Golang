package main

import (
	"fmt"
	"time"
)

func main() {
	//o go na frente de uma função é como se você dissesse para o programa:
	//execute essa função, mas não espere ela terminar para continuar o programa
	go escrever("Olá mundo!")
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}

}
