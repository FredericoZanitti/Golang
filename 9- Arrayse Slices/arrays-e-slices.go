package main

import (
	"fmt"
	"reflect"
)

func main() {
	//ARRAYS

	var array1 [5]string
	fmt.Println(array1)
	array1[0] = "Posição 1"

	array2 := [5]string{"pos1", "pos2", "pos3", "pos4", "pos5"}

	fmt.Println(array1, array2)

	//fixar o tamanho do array de acordo com a quantidade de itens passados
	array3 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(array3)

	//SLICES
	slice := []int{10, 11, 22, 44, 45, 8, 9, 3}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 55)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	//arrays internos
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //tamanho
	fmt.Println(cap(slice3)) //capacidade
	//criou um array de 15 posições e o slice retornou as 10 primeiras posições desse array

	slice4 := make([]float32, 10, 11)
	fmt.Println(slice4)
	slice4 = append(slice4, 5)
	slice4 = append(slice4, 6)
	fmt.Println(len(slice4)) //tamanho
	fmt.Println(cap(slice4)) //capacidade
	//a capacidade é 11, mas quando estouramos a sua capacidade, o Go cria um novo array
	//dobrando sua capacidade

	//resumindo
	//array => uma lista COM tamanho fixo
	//slice => uma lista SEM tamanho fixo

}
