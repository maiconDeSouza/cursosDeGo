package main

import "fmt"

func somar(n1, n2 int8) int8 {
	return n1 + n2
}

func calc(n1, n2 int16) (int16, float32) {
	mult := n1 * n2
	div := float32(n1) / float32(n2)

	return mult, div
}

func main() {
	soma := somar(3, 5)

	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println("Função F")
	}

	f("Olá função F")

	mult, div := calc(23, 11)

	fmt.Println(mult, div)

}
