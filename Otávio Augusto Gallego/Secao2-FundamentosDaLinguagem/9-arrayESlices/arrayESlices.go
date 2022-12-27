package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	var array2 = [3]string{"Posição 1", "Posição 2", "Posição 3"}

	fmt.Println(array2)

	slice := []int{2, 3}

	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array2))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 18, 19, 21)

	fmt.Println(slice)

	slice2 := slice[1:4]

	fmt.Println(slice2)
	slice2[0] = 202
	fmt.Println(slice)
	fmt.Println(slice2)

	slice3 := array2[0:2]
	slice3[0] = "Posição Nova"
	slice3 = append(slice3, "New New")

	// array2 = append(array2, "nnnnn")

	fmt.Println(slice3)
	fmt.Println(array2)

	//Todo slice aponta para um Array

	//array interno

	slice4 := make([]float32, 0, 5)

	slice4 = append(slice4, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7)

	fmt.Println(slice4, cap(slice4))

}
