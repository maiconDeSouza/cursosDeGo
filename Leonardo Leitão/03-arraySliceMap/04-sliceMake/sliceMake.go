package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(25) + 1
}

func main() {
	s := make([]int, 0, 20)

	fmt.Println(s, len(s), cap(s))

	s = append(s, 23, 5, 18)

	fmt.Println(s, len(s), cap(s))

	parada := cap(s) + 1

	fmt.Println(parada)

	for len(s) != parada {
		s = append(s, numeroAleatorio())
	}

	fmt.Println(s, len(s), cap(s))

}
