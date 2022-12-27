package main

import "fmt"

func notaParaConceito(nota float64) string {
	switch {
	case nota >= 9 && nota < 11:
		return "A"
	case nota >= 7 && nota < 9:
		return "B"
	case nota >= 5 && nota < 7:
		return "C"
	case nota >= 3 && nota < 5:
		return "D"
	case nota >= 0 && nota < 3:
		return "E"
	default:
		return "Nota invalida!!!"
	}
}

func main() {
	fmt.Println(notaParaConceito(3))
}
