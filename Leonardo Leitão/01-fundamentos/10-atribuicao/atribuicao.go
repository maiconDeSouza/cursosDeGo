package main

import (
	"fmt"
	"math"
)

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2

	fmt.Println(i)

	var maxI uint32 = math.MaxUint32

	fmt.Println(maxI)

	maxI++

	fmt.Println(maxI)

	x, y := 1, 2
	x, y = y, x

	fmt.Println(x, y)
}
