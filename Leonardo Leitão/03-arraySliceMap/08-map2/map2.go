package main

import "fmt"

func main() {
	funESalarios := map[string]float64{
		"Jose João":      11325,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.00,
	}

	funESalarios["Rafael Filho"] = 1350.27

	fmt.Println(funESalarios)
	delete(funESalarios, "Inexistente")

	for nome, salario := range funESalarios {
		fmt.Println(nome, salario)
	}
}
