package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//fugi totalmente da aula kkkk

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(25) + 1
}

func getValueSlice(arrayJS []int, numero int) bool {
	isValid := false

	for _, value := range arrayJS {
		if value == numero {
			isValid = true
		}
	}

	return isValid
}

func main() {
	var arrayJS []int

	isValid := true

	for isValid {
		n := numeroAleatorio()

		haveNumber := getValueSlice(arrayJS, n)

		if !haveNumber {
			arrayJS = append(arrayJS, n)

			if length := len(arrayJS); length == 15 {
				isValid = false
			}
		}
	}

	sort.Ints(arrayJS)

	fmt.Println(arrayJS)
}
