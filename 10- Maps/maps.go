package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Fred",
		"sobrenome": "Zanitti",
	}

	fmt.Println(usuario)
	//para acessar o campo
	fmt.Println(usuario["nome"])

	//aninhando maps
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":   "Engenharia de Dados",
			"campus": "Campos I",
		},
	}

	fmt.Println(usuario2)
	//se quiser deletar uma chave, use delete passando o map e a chave que deseja deletar
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	//adicionando chaves
	usuario2["signo"] = map[string]string{
		"nome": "Sagitário",
	}
	fmt.Println(usuario2)

}
