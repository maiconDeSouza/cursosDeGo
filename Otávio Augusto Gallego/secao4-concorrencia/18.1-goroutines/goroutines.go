package main

import (
	"fmt"
	"time"
)

func escrevendo(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	//Concorrência != Paralelismo
	go escrevendo("Olá Mundo!!!!") //goroutine
	escrevendo("Programando em Go!!!!")
}
