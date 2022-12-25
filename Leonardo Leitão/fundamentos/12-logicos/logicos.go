package main

import "fmt"

func compras(trab1, trab2 bool) (bool, bool, bool) {
	comprarTv50 := trab1 && trab2
	comprarTv32 := trab1 != trab2
	tomarSorverte := trab1 || trab2

	return comprarTv50, comprarTv32, tomarSorverte
}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("Comprar TV de 50: %v\n Comprar Tv de 32: %v\n Tomar sorvete: %v\n Saudavel: %v", tv50, tv32, sorvete, !sorvete)
}
