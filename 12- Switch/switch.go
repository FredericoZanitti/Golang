package main

import "fmt"

//maneira de utilização I
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feria"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

//maneira de utilização II
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feria"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

//maneira de utilização III
func diaDaSemana3(numero int) string {
	var diaSemana string

	switch {
	case numero == 1:
		diaSemana = "Domingo"
	case numero == 2:
		diaSemana = "Segunda-Feira"
	case numero == 3:
		diaSemana = "Terça-Feria"
	case numero == 4:
		diaSemana = "Quarta-Feira"
	case numero == 5:
		diaSemana = "Quinta-Feira"
	case numero == 6:
		diaSemana = "Sexta-Feira"
	case numero == 7:
		diaSemana = "Sábado"
	default:
		diaSemana = "Dia da Semana Inválido"
	}

	return diaSemana
}

func main() {
	dia := diaDaSemana(4)
	fmt.Println(dia)

	dia2 := diaDaSemana2(5)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(3)
	fmt.Println(dia3)
}
