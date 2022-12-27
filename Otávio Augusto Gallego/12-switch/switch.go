package main

import "fmt"

func diaDaSemana1(numero int8) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabádo"
	default:
		return "Digite um número valido"
	}
}

func diaDaSemana2(numero int8) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	case numero == 3:
		return "Terça-Feira"
	case numero == 4:
		return "Quarta-Feira"
	case numero == 5:
		return "Quinta-Feira"
	case numero == 6:
		return "Sexta-Feira"
	case numero == 7:
		return "Sabádo"
	default:
		return "Digite um número valido"
	}
}

func main() {
	dia := diaDaSemana1(1)

	fmt.Println(dia)

	dia2 := diaDaSemana2(3)

	fmt.Println(dia2)
}
