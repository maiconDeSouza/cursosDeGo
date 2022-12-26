package main

import (
	"fmt"
)

type Usuario struct {
	nome     string
	idade    int8
	endereco Endereco
}

type Endereco struct {
	rua    string
	numero int8
}

func main() {
	end := Endereco{
		rua:    "Rua One",
		numero: 123,
	}
	usuario := Usuario{
		nome:     "Mcn",
		idade:    23,
		endereco: end,
	}

	fmt.Println(usuario)
}
