package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 { //Lembra o While
		fmt.Println(i)
		i++
	}

	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	for {
		//laço infinito

		fmt.Println("Loop...")
		time.Sleep((time.Second * 5))
	}
}
