package main

// Nao existe
func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	println(obterResultado(6.2))
}
