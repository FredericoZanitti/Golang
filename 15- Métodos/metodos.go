package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// metodo associado a uma estrutura
func (u usuario) ImprimirDados() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//usando o ponteiro vc referencia o atributo "fora" do metodo e a atualização acontece
func (u *usuario) fazAniversario() {
	u.idade++
	fmt.Println(u.idade)
}

//nesse caso como não é um ponteiro (u usuario) a atualização acontece somente dentro da função
func (u usuario) fazAniversario2() {
	u.idade++
	fmt.Println(u.idade)
}

func main() {
	usuario1 := usuario{"Fred", 45}
	usuario1.ImprimirDados()
	fmt.Println(usuario1.maiorDeIdade())
	usuario1.fazAniversario()
	fmt.Println(usuario1.idade)
	usuario1.fazAniversario2()
	fmt.Println(usuario1.idade)

}
