package main

import "fmt"

type generico interface{}

func qualquerCoisa(g ...generico) {
	fmt.Println(g)
}

func main() {
	qualquerCoisa("Oi")
	qualquerCoisa(23)
	qualquerCoisa(true, "oi", 23)

}
