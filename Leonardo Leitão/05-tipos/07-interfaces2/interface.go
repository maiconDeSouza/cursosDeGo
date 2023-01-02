package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
}

func main() {
	f200 := ferrari{
		modelo:          "f200",
		turboLigado:     false,
		velocidadeAtual: 120,
	}

	var f40 esportivo = &ferrari{"f40", false, 121}

	f200.ligarTurbo()
	f40.ligarTurbo()

	fmt.Println(f200)
	fmt.Println(f40)
}
