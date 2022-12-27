package main

import (
	"app-linha-de-comando/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	application := app.Gerar()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
