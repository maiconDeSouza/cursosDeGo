package main

import "fmt"

//panic mata o programa

func div(x, y int) int {
	defer contraPanic()
	return x / y
}

func contraPanic() {

	if r := recover(); r != nil {
		fmt.Println("O Erro cantou", r)
	}

}

func main() {
	fmt.Println(div(1, 0))
	fmt.Print(div(6, 3))
}
