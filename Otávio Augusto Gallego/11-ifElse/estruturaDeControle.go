package main

import "fmt"

func main() {
	const numero = -1

	if numero >= 5 {
		fmt.Printf("%v é maior ou igual a 5\n", numero)
	} else {
		fmt.Printf("%v é menor que 5\n", numero)
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Printf("%v é maior que 0", outroNumero)
	} else {
		fmt.Printf("%v é menor ou igual a 0", outroNumero)
	}
}
