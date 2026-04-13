package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")

	fmt.Println("Nova")
	fmt.Println("Linha")

	x := 3.1415
	// fmt.Println("O valor de x é ", )

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)
	fmt.Printf("O valor de x é %.2f\n", x)

	a := 1
	b := 1.99
	c := false
	d := "opa"
	fmt.Printf("O valor de a é %d, b é %.2f, c é %t e d é %s\n", a, b, c, d)

}
