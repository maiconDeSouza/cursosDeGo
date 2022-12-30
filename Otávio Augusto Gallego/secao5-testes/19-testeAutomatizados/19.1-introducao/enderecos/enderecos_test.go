package enderecos

import (
	"testing"
)

type structParaTeste struct {
	valorPassado   string
	valorRetornado string
}

var valoresParaTeste = []structParaTeste{
	{"Rua ABC", "Rua"},
	{"RUA ABC", "Rua"},
	{"Avenida ABC", "Avenida"},
	{"Estrada ABC", "Estrada"},
	{"Caminho ABC", "Inválido!"},
	{"Alameda ABC", "Alameda"},
	{" ABC", "Inválido!"},
	{"Praça ABC", "Inválido!"},
	{"Rodovia ABC", "Rodovia"},
	{"Parque ABC", "Inválido!"},
	{"Garagem ABC", "Inválido!"},
}

func TestTipoDeEnderco(t *testing.T) {
	for _, valor := range valoresParaTeste {
		retorno := TipoDeEndereco(valor.valorPassado)

		if retorno != valor.valorRetornado {
			t.Errorf("Para o Valor %v o esperado era %v", valor.valorPassado, valor.valorRetornado)
		}
	}
}
