package main

func notaParaConceito(nota float64) string {
	switch {
	case nota >= 9.0:
		return "A"
	case nota >= 7.0:
		return "B"
	case nota >= 5.0:
		return "C"
	case nota >= 3.0:
		return "D"
	default:
		return "E"
	}
}

func main() {
	println(notaParaConceito(9.5))
	println(notaParaConceito(8.0))
	println(notaParaConceito(6.0))
	println(notaParaConceito(4.0))
	println(notaParaConceito(2.0))
}
