package main

import "fmt"

func main() {
	// var aprovados map[int]string

	aprovados := make(map[int]string)
	aprovados[1234567890] = "Maria"
	aprovados[1234567891] = "Pedro"
	aprovados[1234567892] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[987654342])
	delete(aprovados, 1234567891)
	fmt.Println(aprovados)

}
