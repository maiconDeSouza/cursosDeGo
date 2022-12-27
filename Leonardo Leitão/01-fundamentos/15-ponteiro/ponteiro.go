package main

import "fmt"

func main() {
	i := 1

	var p1 *int = nil

	p1 = &i

	var p2 *int = &i

	*p1++
	i++

	fmt.Println(i, *p1, *p2, &i, p1, p2)
	fmt.Println(*p2 == *p1)

}
