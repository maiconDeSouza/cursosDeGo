package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return fmt.Sprintf("%s %s", p.nome, p.sobrenome)
}

func (p produto) toString() string {
	return fmt.Sprintf("%s - R$%.2f", p.nome, p.preco)

}

func print(x imprimivel) {
	fmt.Println(x.toString())
}

func main() {
	person := pessoa{
		nome:      "Elon",
		sobrenome: "Musk",
	}

	caixaDeSom := produto{
		nome:  "JBL",
		preco: 232.23,
	}

	print(person)
	print(caixaDeSom)

	var coisa imprimivel = pessoa{"Bill", "Gates"}

	print(coisa)

	coisa = produto{"Caneta Azul", 4.87}

	print(coisa)

}
