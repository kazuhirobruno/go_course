package main

func main() {
	numeros := [...]int{1, 2, 3, 4, 5}
	for i, numero := range numeros {
		println(i+1, numero)
	}

	for _, num := range numeros {
		println(num)
	}
}
