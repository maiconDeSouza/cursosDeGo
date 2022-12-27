package main

import "fmt"

func clousure() func() {
	text := "Dentro da Função Closure"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func umNumero(a int) func() {
	umTexto := func() {
		fmt.Printf("%v - %v", "Text", a)

	}

	return umTexto
}

func main() {
	text := "Função dentro do Main"
	fmt.Printf("%T\n", text)

	funcaoNova := clousure()
	funcaoNova()

	p1 := umNumero(23)
	p1()

}

//JS mais flexivel no clousure
