package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	segundo = p1
	primeiro = p2

	return //retorno limpo
}

func main() {
	segundo, primeiro := trocar(23, 32)

	fmt.Println(primeiro, segundo)
}
