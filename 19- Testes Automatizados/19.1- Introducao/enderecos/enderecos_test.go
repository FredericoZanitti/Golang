package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeteste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paranaguá", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Estrada Sem Número", "Estrada"},
		{"Praça Sete", "Tipo Inválido"},
		{"RUA JARAGUÁ", "Rua"},
		{"BECO REBOLÇAS", "Beco"},
		{"rodovia higinópolis", "Rodovia"},
	}

	for _, cenario := range cenariosDeteste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
