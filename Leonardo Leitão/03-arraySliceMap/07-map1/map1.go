package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[123] = "Maria"
	aprovados[456] = "Pedro"
	aprovados[789] = "Ana"

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados)

	delete(aprovados, 456)

	fmt.Println(aprovados)

}
