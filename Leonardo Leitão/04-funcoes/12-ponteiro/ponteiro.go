package main

import "fmt"

var sl = []int{1, 2, 3}

func main() {
	n := 17

	inc1(n)

	inc2(&n)

	fmt.Println(n)

	slFunc(sl)
	fmt.Println(sl)

}

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func slFunc(sl []int) {
	sl = append(sl, 23, 45)
	fmt.Println(sl)
}
