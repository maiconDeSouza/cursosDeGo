package main

import (
	"fmt"
	"time"
)

func main() {
	ch := escrever("Programando em Go!!!!")

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func escrever(texto string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Second * 2)
		}
	}()

	return ch
}
