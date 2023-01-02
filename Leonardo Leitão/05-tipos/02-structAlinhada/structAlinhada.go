package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (p pedido) totalPedido() float64 {
	var total float64

	for _, item := range p.itens {
		total += (item.preco * float64(item.qtde))
	}

	return total
}

func main() {
	item1 := item{
		produtoID: 1,
		qtde:      3,
		preco:     1.21,
	}

	item2 := item{
		produtoID: 2,
		qtde:      1,
		preco:     1.5,
	}

	item3 := item{
		produtoID: 3,
		qtde:      3,
		preco:     4.23,
	}

	cliente1 := pedido{
		userID: 123,
		itens:  []item{item1, item2, item3},
	}

	fmt.Println(cliente1.totalPedido())

}
