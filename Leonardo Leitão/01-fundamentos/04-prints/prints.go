package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("O Valor de x depois de usar o Sprint" + xs) //aqui posso usar o + de concatenação

	fmt.Println("O Valor de X sem usar a concatenação", x)

	fmt.Printf("O Valor de x usando o PrintF é %.2f \n", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("%d - %f - %t - %s \n", a, b, c, d)
	fmt.Printf("%v - %v - %v - %v", a, b, c, d)

}
