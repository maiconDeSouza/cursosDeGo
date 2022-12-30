package main

import (
	"fmt"
	"introducaoATeste/enderecos"
)

func main() {
	retorno := enderecos.TipoDeEndereco("Avenida Paulista")

	fmt.Println(retorno)
}
