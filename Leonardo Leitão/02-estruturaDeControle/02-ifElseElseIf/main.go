package main

import (
	"errors"
	"fmt"
)

func notaParaConceito(nota int) {
	if err := errors.New("Nota Invalida!"); nota > 10 || nota < 0 {
		fmt.Println(err)
		return
	}

	if nota >= 9 {
		fmt.Println("A")
	} else if nota >= 8 {
		fmt.Println("B")
	} else if nota >= 5 {
		fmt.Println("C")
	} else if nota >= 3 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}
}

func main() {
	notaParaConceito(2)
}
