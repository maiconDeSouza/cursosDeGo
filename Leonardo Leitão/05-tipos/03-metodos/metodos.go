package main

import (
	"fmt"
	"strings"
)

func main() {
	person := pessoa{
		name:     "Oliver",
		lastname: "Tsubasa",
	}

	person.getFullName()
	person.upName("Carlos Santana")
	person.getFullName()

}

type pessoa struct {
	name     string
	lastname string
}

func (p pessoa) getFullName() {
	fmt.Printf("%s %s\n", p.name, p.lastname)
}

func (p *pessoa) upName(fullName string) {
	sliceFullName := strings.Fields(fullName)

	p.name = sliceFullName[0]
	p.lastname = sliceFullName[1]
}
