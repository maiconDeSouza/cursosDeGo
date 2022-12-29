package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	a := soma

	fmt.Println(a(2, 3))
}
