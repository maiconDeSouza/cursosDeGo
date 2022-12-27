package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {

	posicao := fibonacci(40)

	fmt.Println(posicao)

	// 1 1 2 3 5 8 13 21 34 55
}
