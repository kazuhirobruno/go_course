package main

import "fmt"

func int1(n int) {
	n++
}

func int2(n *int) {
	*n++
}

func main() {
	n := 1
	int1(n)
	fmt.Println("int1:", n)
	int2(&n)
	fmt.Println("int2:", n)
}
