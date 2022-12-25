package main

import (
	"errors"
	"fmt"
)

func somaResultEven(x int16, y int16) (int16, error) {
	result := x + y
	if result%2 != 0 {
		return result, errors.New("Este número é Impar!!!")
	}
	return result, nil
}

func imprimir(result int16, err error) string {
	if err != nil {
		return fmt.Sprintf("%v é um valor invalido!", result)
	}

	return fmt.Sprintf("%v é valor valido", result)
}

func main() {
	teste1 := imprimir(somaResultEven(23, 35))

	teste2 := imprimir(somaResultEven(3, 4))

	fmt.Println(teste1)
	fmt.Println(teste2)

}
