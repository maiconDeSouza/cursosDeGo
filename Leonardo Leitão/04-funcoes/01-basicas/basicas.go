package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5(n1, n2 int) (int, int) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	f1()

	f2("Eu", "Programo em Go!!!")

	retornoF3 := f3()
	fmt.Println(retornoF3)

	retornoF4 := f4("Go", "Uma ótima linguagem")
	fmt.Println(retornoF4)

	soma, sub := f5(250, 123)
	fmt.Println(soma, sub, soma+sub)
}
