package main

import "fmt"

func soma(numeros ...int) (resultado int) {
	for _, valor := range numeros {
		resultado += valor
	}
	return
}

func main() {
	resultado := soma(23, 45, 98, 54, 3, 89, 90, 2, 18)

	fmt.Println(resultado)
}
