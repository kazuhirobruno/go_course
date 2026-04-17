package main

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}
	println(Distancia(p1, p2)) // Deve imprimir 5
	println(Catetos(p1, p2))
}
