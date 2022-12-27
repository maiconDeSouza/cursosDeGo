package main

import "fmt"

type cadastroDeBalada struct {
	nome  string
	idade int8
}

func (c cadastroDeBalada) podeEntrar() {
	if c.idade >= 18 {
		fmt.Printf("Olá %v você tem %v e pode entrar!\n", c.nome, c.idade)
	} else {
		fmt.Printf("Olá %v você tem %v e não pode entrar!\n", c.nome, c.idade)
	}
}

func main() {
	p1 := cadastroDeBalada{"Jorge", 23}
	p2 := cadastroDeBalada{"Rita", 35}
	p3 := cadastroDeBalada{"Ricardo", 17}

	p3.idade = 55

	p1.podeEntrar()
	p2.podeEntrar()
	p3.podeEntrar()
}
