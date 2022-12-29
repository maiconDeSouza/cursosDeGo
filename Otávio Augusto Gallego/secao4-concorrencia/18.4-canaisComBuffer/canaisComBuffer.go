package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Oiii"
	ch <- "Programando em Go"

	message := <-ch
	message2 := <-ch

	fmt.Println(message)
	fmt.Println(message2)
}
