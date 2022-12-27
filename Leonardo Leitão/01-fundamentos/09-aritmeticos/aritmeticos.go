package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Divisão =>", float64(a)/float64(b))
	fmt.Println("Módulo =>", a%b)

	//bitwise
	fmt.Println("E =>", a&b)
	fmt.Println("Ou =>", a|b)
	fmt.Println("Xor =>", a^b)

	c := 5.0
	d := 7.0

	//outras operações usando math...
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Maior =>", math.Max(c, d))

}
