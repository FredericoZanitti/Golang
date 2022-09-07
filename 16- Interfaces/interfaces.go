package main

import (
	"fmt"
	"math"
)

// a interface tem a assinatura do metodo e todos os métodos que desejarem usar essa interface
// devem ter um método associado com essa mesma assinatura "area() float64", ou seja, um método
// area que não recebe parâmetros e que retorna um float64
type forma interface {
	area() float64
}

type retangulo struct {
	base   float64
	altura float64
}

// para usar a interface basta associar um método identico ao da interface no método retangulo
func (r retangulo) area() float64 {
	return r.base * r.altura
}

type circulo struct {
	raio float64
}

// da mesma forma, com o círculo, implementar um método identico ao da interface
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	r := retangulo{5.5, 6.9}
	fmt.Printf("Área do Retângulo: %0.2f\n", r.area())

	c := circulo{4.5}
	fmt.Printf("Área do Círculo: %0.2f\n", c.area())
}
