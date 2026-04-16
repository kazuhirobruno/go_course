package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type Ferrari struct {
	modelo          string
	turboLigada     bool
	velocidadeAtual int
}

func (f *Ferrari) ligarTurbo() {
	f.turboLigada = true
}

func main() {
	carro1 := &Ferrari{modelo: "F40", turboLigada: false, velocidadeAtual: 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &Ferrari{modelo: "F50", turboLigada: false, velocidadeAtual: 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
