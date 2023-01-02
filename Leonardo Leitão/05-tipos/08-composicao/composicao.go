package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type mercedes struct {
	modelo string
}

func (m mercedes) ligarTurbo() {
	fmt.Println("Ligando o Turbo do Modelo Mercedes", m.modelo)
}

func (m mercedes) fazerBaliza() {
	fmt.Println("Fazendo a baliza automaticamente do Modelo Mercedes", m.modelo)
}

func main() {
	var c200 luxuoso = mercedes{"c200"}
	var c300 esportivo = mercedes{"c300"}
	var c500 esportivoLuxuoso = mercedes{"c500"}

	c200.fazerBaliza()
	c300.ligarTurbo()

	c500.fazerBaliza()
	c500.ligarTurbo()
}
