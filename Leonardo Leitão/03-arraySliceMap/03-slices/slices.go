package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(reflect.TypeOf(s1))
	fmt.Println(len(s1), cap(s1))

	s1 = append(s1, 23)

	fmt.Println(len(s1), cap(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]
	s3 := a2[1:3]

	s2 = append(s2, 21)

	fmt.Println(s2, s3)
	s2[0] = 35

	fmt.Println(s2, s3)

	fmt.Println(s2, len(s2), cap(s2))

}
