package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		formataFatorial(5)
		wg.Done()
	}()

	go func() {
		formataFatorial(-4)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Fim!")
}

func formataFatorial(n int) {
	result, err := fatorial(n)
	if err != nil {
		fmt.Printf("Ocorreu um erro -> %v\n", err)
		return
	}
	fmt.Printf("O fatorial de %v é %v\n", n, result)

}

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}
