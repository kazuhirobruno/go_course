package main

import "strings"

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) mudarNome(nomeCompleto string) {
	partes := strings.Split(nomeCompleto, " ")
	p.nome = partes[0]
	p.sobrenome = partes[1]
}

func main() {
	p1 := pessoa{"João", "Silva"}
	println(p1.nomeCompleto())
	p1.mudarNome("Maria Oliveira")
	println(p1.nomeCompleto())
}
