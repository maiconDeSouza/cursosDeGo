package main

import "fmt"

var aprovados = []string{"Maria", "Pedro", "Rafael", "Guilherme"}

func main() {
	imprimirAprovados(aprovados...)
}

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Aprovados!!!!")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
