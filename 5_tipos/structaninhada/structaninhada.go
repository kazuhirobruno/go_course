package main

import "fmt"

type item struct {
	productId int
	quantity  int
	preco     float64
}

type pedido struct {
	userid int
	items  []item
}

func (p pedido) total() float64 {
	total := 0.0
	for _, item := range p.items {
		total += float64(item.quantity) * item.preco
	}
	return total
}

func main() {
	p1 := item{1, 2, 3000.00}
	p2 := item{2, 1, 1500.00}
	p3 := item{3, 3, 500.00}

	p := pedido{userid: 1, items: []item{p1, p2, p3}}
	fmt.Printf("Total do pedido: %.2f\n", p.total())
}
