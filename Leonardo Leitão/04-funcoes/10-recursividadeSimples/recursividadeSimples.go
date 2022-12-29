package main

import "fmt"

func main() {
	formataFatorial(65)
}

func formataFatorial(n uint64) {
	result := fatorial(n)
	fmt.Printf("O fatorial de %v é %v\n", n, result)

}

func fatorial(n uint64) uint64 {
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior
	}
}
