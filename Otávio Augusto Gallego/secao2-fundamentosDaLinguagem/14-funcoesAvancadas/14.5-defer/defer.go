package main

import "fmt"

func funcao1() {
	fmt.Println("Executando Função 1")
}

func funcao2() {
	fmt.Println("Executando Função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media está sendo Calculada...")
	fmt.Println("Função de aprovação")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}

	return false
}

func main() {
	// defer funcao1()
	// funcao2()

	fmt.Println(alunoAprovado(8, 6))
}
