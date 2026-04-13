package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Resto da divisão:", a%b)

	fmt.Println("E: ", a&b)
	fmt.Println("OU: ", a|b)
	fmt.Println("XOR: ", a^b)
	fmt.Println("NOT: ", ^a)

	c := 3.0
	d := 2.0
	fmt.Println("Maior: ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor: ", math.Min(c, d))
	fmt.Println("Potência: ", math.Pow(c, d))
	fmt.Println("Raiz quadrada: ", math.Sqrt(c))

}
