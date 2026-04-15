package main

func notaParaConceito(nota float64) string {
	if nota >= 9 {
		return "A"
	} else if nota >= 7 {
		return "B"
	} else if nota >= 5 {
		return "C"
	} else if nota >= 3 {
		return "D"
	} else {
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
