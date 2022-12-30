package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint8  `json:"idade"`
}

func main() {
	dog1 := cachorro{"Dante", "Pug", 2}

	fmt.Println(dog1)

	dogJson, err := json.Marshal(dog1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dogJson))
}
