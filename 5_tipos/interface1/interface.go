package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	nome  string
	preco float64
}

func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	return p.nome + " custa R$" + fmt.Sprintf("%.2f", p.preco)
}

func imprimir(i imprimivel) {
	fmt.Println(i.toString())
}

func main() {
	p1 := pessoa{nome: "João", sobrenome: "Silva"}
	pr1 := produto{nome: "Notebook", preco: 2500.00}

	imprimir(p1)
	imprimir(pr1)
}
