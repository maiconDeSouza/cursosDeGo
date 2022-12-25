package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Strings", "Abacaxi" == "Abacaxi")
	fmt.Println("!=", 5 != 8)
	fmt.Println("<", 5 < 8)
	fmt.Println(">", 5 > 8)
	//>= <=

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)

	fmt.Println("Datas", d1 == d2)
	fmt.Println("Datas", d1.Equal(d2))

	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"João"}
	p2 := Pessoa{"João"}

	fmt.Println(p1)
	fmt.Println(p1 == p2)
}
