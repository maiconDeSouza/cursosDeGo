package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

// Registra
func main() {
	fmt.Println("Rodando do pacote main")

	auxiliar.Escrevendo()

	fmt.Println(checkmail.ValidateFormat("mcnhotmail.com"))
}
