package main

import (
	"fmt"
	"math"
)

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// Método: Função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	x := p.preco * (1 - p.desconto)

	return math.Round(x*100) / 100
}

func main() {
	produto1 := produto{
		nome:     "Lapis",
		preco:    1.87,
		desconto: 0.05,
	}

	fmt.Println(produto1)
	fmt.Println(produto1.precoComDesconto())
}
