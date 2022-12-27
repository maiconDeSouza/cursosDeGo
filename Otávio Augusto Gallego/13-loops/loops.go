package main

import (
	"fmt"
	"time"
)

type person struct {
	name string
	age  int8
}

var lista []person

func criaListaDePerson(name string, age int8) {
	lista = append(lista, person{
		name: name,
		age:  age,
	})
}

func main() {
	i := 0

	for i < 10 {
		i++

		fmt.Println(i)

		time.Sleep(time.Second)
	}

	for jx := 10; jx > 0; jx-- {
		fmt.Println(jx)

		time.Sleep(time.Second)
	}

	nome := [3]string{"Lucas", "Dante", "Taluzinha"}

	for i, item := range nome {
		fmt.Println(i, item)
	}

	criaListaDePerson("Dante", 2)
	criaListaDePerson("Talu", 13)
	criaListaDePerson("Felipe", 11)

	fmt.Println(lista)

}
