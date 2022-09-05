package main

import (
	"errors"
	"fmt"
)

func main() {
	//tipos de inteiros: int8, int16, int32, int64
	var numero int32 = 191914
	fmt.Println(numero)

	//tipos de num reais: float32, float64
	var numeroReal float32 = 15.252
	fmt.Println(numeroReal)

	//string - não existe CHAR no go
	var str string = "Texto qualquer"
	fmt.Println(str)

	//se não iniciar uma variável, ela será
	//0 para int/float
	//string em branco para string
	//false para boolean

	//tipo booleano
	var booleano bool = true
	fmt.Println(booleano)

	//Erro é um tipo
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
