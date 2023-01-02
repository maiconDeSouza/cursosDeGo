package main

import "fmt"

type carro struct {
	nome       string
	velocidade int
}

func (c *carro) acelerar() {
	c.velocidade += 5
}

type ferrari struct {
	carro
}

func (f *ferrari) acelerar() {
	f.velocidade += 10
}

var fusca = carro{
	nome:       "fusca",
	velocidade: 0,
}

var f200 = ferrari{}

func main() {
	fmt.Println(fusca.velocidade)
	fusca.acelerar()
	fmt.Println(fusca.velocidade)

	f200.nome = "f200"
	f200.velocidade = 0

	fmt.Println(f200.velocidade)
	f200.acelerar()
	fmt.Println(f200.velocidade)

}
