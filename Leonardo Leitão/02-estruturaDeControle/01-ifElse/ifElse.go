package main

import (
	"errors"
	"fmt"
)

func imprimirResultado(n int) {
	if err := errors.New("Numero tem que ser no máximo 10"); n > 10 {
		fmt.Println(err)
		return
	}

	if n > 5 {
		fmt.Println("Maior que 5")
	} else {
		fmt.Println("Menor que 5")
	}
}

func main() {
	imprimirResultado(7)
	imprimirResultado(3)
	imprimirResultado(233)
}
