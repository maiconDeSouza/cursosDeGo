package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			ch1 <- "Func 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 1)
			ch2 <- "Func 2"
		}
	}()

	for {

		select {
		case messageFunc1 := <-ch1:
			fmt.Println(messageFunc1)
		case messageFunc2 := <-ch2:
			fmt.Println(messageFunc2)
		}

	}
}
