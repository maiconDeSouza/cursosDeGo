package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println((notaFinal))

	//cuidado...

	//fmt.Println("Teste " + string(97)) //em vez de converter o 97 em uma string ele irá pegar o valor
	//equivalente na tabela ascii

	//int para string
	fmt.Println("teste " + strconv.Itoa(123))

	//string para int
	num1, err1 := strconv.Atoi("123")
	fmt.Println(num1, err1)

	num2, err2 := strconv.Atoi("casa")
	fmt.Println(num2, err2)

	//string para boolean
	bol, err3 := strconv.ParseBool("true") //só funciuonaria 1 e 0 tbm além dos boolean
	fmt.Println(bol, err3)

}
