package main

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9, 10) {
		return "A"
	} else if n.entre(7, 8.99) {
		return "B"
	} else if n.entre(5, 6.99) {
		return "C"
	} else if n.entre(3, 4.99) {
		return "D"
	} else if n.entre(0, 2.99) {
		return "E"
	} else {
		return "Nota inválida"
	}
}

func main() {
	var minhaNota nota = 8.5
	conceito := notaParaConceito(minhaNota)
	println("Minha nota é:", minhaNota, "e o conceito é:", conceito)

	var outraNota nota = 2.5
	outroConceito := notaParaConceito(outraNota)
	println("Outra nota é:", outraNota, "e o conceito é:", outroConceito)
}
