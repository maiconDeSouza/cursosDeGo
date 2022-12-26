package main

import "fmt"

type pessoa struct {
	nome  string
	idade int8
}

type estudante struct {
	pessoa
	curso string
}

func main() {
	p1 := pessoa{"Dante", 2}
	est := estudante{p1, "Medicina"}

	fmt.Println(est)
}
