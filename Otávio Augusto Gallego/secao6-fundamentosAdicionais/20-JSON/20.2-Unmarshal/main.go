package main

import (
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
	cachorroemJSON := `{"nome":"Dante","raca":"Pug","idade":2}`

	c := cachorro{}

	if err := json.Unmarshal([]byte(cachorroemJSON), &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

}
