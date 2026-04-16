package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0
	for _, n := range numeros {
		total += n
	}
	return total / float64(len(numeros))
}

func main() {
	resultado := media(7.5, 8.0, 9.0)
	fmt.Printf("A média é: %.2f\n", resultado)
}
