package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 4 - 2
	divisao := 10 / 2
	multiplicacao := 5 * 7
	restoDaDivisao := 10 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int32 = 25
	somar := numero1 + int16(numero2)
	fmt.Println(somar)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//OPERADORES LÓGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNÁRIOS
	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)
	numero += 15
	fmt.Println(numero)
	numero -= 10
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero *= 3
	fmt.Println(numero)
	numero /= 5
	fmt.Println(numero)
	numero %= 3
	fmt.Println(numero)
}
