package main

import "fmt"

func inverterSinal(numero *int) int {
	*numero = *numero * -1
	return *numero
}

func main() {
	n1 := 25

	inverterSinal(&n1)

	fmt.Println("O Valor da variavel é", n1)
}
