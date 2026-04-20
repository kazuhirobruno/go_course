package main

import "time"

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		println(pessoa, "disse:", texto)
	}
}

func main() {
	fale("Maria", "Olá, tudo bem?", 3)
	fale("João", "Oi, tudo ótimo!", 3)

	go fale("Maria", "Olá, tudo bem?", 3)
	go fale("João", "Oi, tudo ótimo!", 3)

	time.Sleep(time.Second * 5)
	println("Fim")
}
