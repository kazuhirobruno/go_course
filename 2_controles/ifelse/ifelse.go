package main

func imprimirResultado(nota float64) {
	if nota >= 7 {
		println("Aprovado com nota", nota)
	} else {
		println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(8.5)
	imprimirResultado(6.0)
}
