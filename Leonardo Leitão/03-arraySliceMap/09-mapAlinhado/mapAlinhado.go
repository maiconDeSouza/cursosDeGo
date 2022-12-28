package main

import "fmt"

func main() {
	funPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15450.23,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"Jose João": 1111.10,
		},
	}

	fmt.Println(funPorLetra)
}
