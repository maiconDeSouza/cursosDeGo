package main

import (
	"fmt"
	"math"
)

func notaParaConceito(nota int) string {
	switch nota {
	case 9, 10:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota invalida!!!"
	}
}

func arredondarNota(nota float64) int {
	notaArredondada := math.Round(nota)

	return int(notaArredondada)
}

func main() {
	nota := 5.25

	notaRedonda := arredondarNota(nota)

	fmt.Println(notaParaConceito(notaRedonda))

}
