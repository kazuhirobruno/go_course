package main

import "fmt"

func main() {
	funcESalarios := map[string]float64{
		"João":  2000.00,
		"Maria": 3000.00,
		"Pedro": 2500.00,
	}

	funcESalarios["Ana"] = 3500.00
	fmt.Println(funcESalarios)
	delete(funcESalarios, "Pedrow")

	for nome, salario := range funcESalarios {
		fmt.Printf("%s ganha R$ %.2f\n", nome, salario)
	}
}
