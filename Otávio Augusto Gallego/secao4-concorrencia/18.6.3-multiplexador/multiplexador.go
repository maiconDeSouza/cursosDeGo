package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := multiplexar(escrever("Hello Go!"), escrever("Programando em Go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func multiplexar(chEntrada1, chEntrada2 <-chan string) <-chan string {
	chSaida := make(chan string)

	go func() {
		for {
			select {
			case msg := <-chEntrada1:
				chSaida <- msg
			case msg := <-chEntrada2:
				chSaida <- msg
			}
		}
	}()
	return chSaida
}

func escrever(texto string) <-chan string {
	ch := make(chan string)

	go func() {
		for {
			ch <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
		}
	}()

	return ch
}
