package main

import "fmt"

func main() {
	result := media(2.5, 10.00, 8.75, 5.00, 10.00, 9.85)

	fmt.Printf("%.2f \n", result)
}

func media(numeros ...float64) float64 {
	total := 0.0
	quantidade := float64(len(numeros))

	for _, valor := range numeros {
		total += valor
	}

	return total / quantidade
}
