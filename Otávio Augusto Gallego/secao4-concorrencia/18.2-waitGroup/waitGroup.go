package main

import (
	"fmt"
	"sync"
	"time"
)

func escrevendo(text string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrevendo("Função 1")
		waitGroup.Done()
	}()

	go func() {
		escrevendo("Função 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()

	fmt.Println("Fim")
}
