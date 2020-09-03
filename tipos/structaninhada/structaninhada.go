package main

import "fmt"

// struct é um agrupamento de dados
type item struct {
	produtoID  int
	quantidade int
	preco      float64
}

type pedido struct {
	userID int
	items  []item
}

// Método: função  com receiver (receptor)

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.items {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {

	pedido := pedido{
		userID: 1,
		items: []item{ // evite fazer desta forma, prefira explicitr o nome dos campos, traz mais clareza
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{produtoID: 11, // Prefira DESTA FORMA, mais clareza da informação
				quantidade: 100,
				preco:      3.13},
		},
	}

	fmt.Println(pedido, pedido.valorTotal())

	fmt.Printf("Valor Total do pedido %d é R$ %.2f", pedido.userID, pedido.valorTotal())

}
