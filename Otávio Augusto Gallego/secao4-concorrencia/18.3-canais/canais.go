package main

import (
	"fmt"
	"time"
)

func escrevendo(text string, canal chan string) {
	for i := 1; i <= 5; i++ {
		canal <- text
		time.Sleep(time.Second * 2)
	}

	close(canal)
}

func main() {
	canal := make(chan string)

	go escrevendo("Função 1", canal)

	for message := range canal {
		fmt.Println(message)
	}

	fmt.Println("Fim!!!")
}
