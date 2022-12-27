package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	//int8, int16, int32, int64

	var numero int16 = 100
	fmt.Println(numero)

	//uint não aceita negativo

	var numero2 uint16 = 12
	fmt.Println(numero2)

	//alias
	var numero3 rune = 122 //rune = int32 mais usado para tabela ascii
	fmt.Println(reflect.TypeOf(numero3))

	var numero4 byte = 23 //byte = unit8
	fmt.Println(reflect.TypeOf(numero4))

	var numeroReal float32 = 123.23
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 234.23
	fmt.Println(numeroReal2)

	numeroReal3 := 12.12
	fmt.Println(numeroReal3)

	var str string = "Texto"
	var char = 'b'

	fmt.Println(str, char)

	var booleano bool = true
	fmt.Println(booleano)

	var erro error
	fmt.Println(erro)

	var erro2 = errors.New("Novo erro")
	fmt.Println(erro2)

}
