package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	//adicionando um grupo de espera com 2 goroutines
	waitGroup.Add(2)

	//primeira função goroutine que entrará na fila
	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() // é como se desse um -1 no count de goroutines
	}()

	//segunda função goroutine que entrará na fila
	go func() {
		escrever("Programando em Go!")
		waitGroup.Done()
	}()

	waitGroup.Wait() //fala pro programa "esperar" o count do wait chegar em zero
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
