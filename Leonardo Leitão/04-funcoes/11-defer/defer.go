package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	totalDePontos(51, "Joaquinzin")
}

func totalDePontos(pontuacao uint8, nome string) {
	var pontuacaoMax uint8 = 100
	if pontuacao > pontuacaoMax {
		fmt.Println("Nota inválida!")
		return
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	pontuacaoDoUltimoAprovado := r.Intn(int(pontuacaoMax)) + 1

	defer fmt.Println("Fechando consulta!!!!")
	defer fmt.Printf("A pontuação do último aprovado foi %v\n", pontuacaoDoUltimoAprovado)

	if pontuacao > uint8(pontuacaoDoUltimoAprovado) {
		fmt.Printf("Olá %v você foi aprovado com a nota %v\n", nome, pontuacao)
		return
	} else {
		fmt.Printf("Olá %v você foi reprovado com a nota %v\n", nome, pontuacao)
		return
	}

}
