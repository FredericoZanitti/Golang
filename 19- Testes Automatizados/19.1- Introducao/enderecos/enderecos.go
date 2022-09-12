package enderecos

import "strings"

// TipoDeEndereco para verificar a validade do endereço passado
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "beco", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"

	// rodar teste em todas as pastas do pacote
	// digite go test ./...

	// rodar teste com mais especificações
	// digite go test -v

	// saber quantos % das funções estão sendo cobertas pelo teste
	// digite go test -cover

	// criar um arquivo txt com o resultado do teste
	// go test --coverprofile cobertura.txt

	// saber exatamente a linha/função onde o teste não cobriu 100%
	// digite go tool cover --func=cobertura.txt
	//        go tool cover --html=cobertura.txt

}
