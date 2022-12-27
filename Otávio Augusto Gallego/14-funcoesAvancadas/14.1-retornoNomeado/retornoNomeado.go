package main

import "fmt"

func calc(n1, n2 int8) (soma int8, subtracao int8) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma, sub := calc(5, 3)

	fmt.Println(soma, sub)
}
