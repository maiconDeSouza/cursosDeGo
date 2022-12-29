package main

import "fmt"

func main() {
	result := exec(multiplicacao, 3, 4)

	fmt.Println(result)
}

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) string {
	retorno := funcao(p1, p2)

	return fmt.Sprintf("O Resultado é %v", retorno)
}
