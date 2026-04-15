package main

import (
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i < 5 {
		println("O número é menor que 5:", i)
	} else {
		println("O número é maior ou igual a 5:", i)
	}
}
