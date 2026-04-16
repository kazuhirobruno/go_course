package main

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	p1 := produto{"Notebook", 3000.00, 0.10}
	p2 := produto{"Smartphone", 1500.00, 0.15}
	println(p1.nome, p1.preco, p1.precoComDesconto())
	println(p2.nome, p2.preco, p2.precoComDesconto())
}
