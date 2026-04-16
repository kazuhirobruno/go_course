package main

func main() {
	funcPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela": 1000.00,
			"Guga":     2000.00,
		},
		"J": {
			"João":  3000.00,
			"Jorge": 4000.00,
			"José":  5000.00,
		},
		"P": {
			"Paulo":  6000.00,
			"Pedro":  7000.00,
			"Paula":  8000.00,
			"Pamela": 9000.00,
		},
	}

	delete(funcPorLetra, "P")

	for letra, funcs := range funcPorLetra {
		println("Letra:", letra)
		for nome, salario := range funcs {
			println("  ", nome, salario)
		}
	}
}
