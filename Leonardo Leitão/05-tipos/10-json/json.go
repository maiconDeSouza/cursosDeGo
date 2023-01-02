package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func TesteErro() error {
	err := errors.New("Deu erro")
	return err
}

func main() {
	//struct para json
	p1 := produto{1, "Notebook", 1899.99, []string{"Promocao", "Eletronico"}}
	p1JSON, errM := json.MarshalIndent(p1, "", "  ")
	Panic(errM)

	fmt.Println(string(p1JSON))

	//json para struct
	jsonString := `{"id":2, "nome":"Caneta", "tags":["azul"]}`
	var p2 produto
	errU := json.Unmarshal([]byte(jsonString), &p2)
	Panic(errU)
	fmt.Println(p2)

	errN := TesteErro()
	Panic(errN)

}
