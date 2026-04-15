package main

import "time"

func main() {
	i := 1

	for i <= 10 {
		println(i)
		i++
	}

	for i := 0; i <= 20; i++ {
		println(i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			println(i, "Par")
		} else {
			println(i, "Impar")
		}
	}

	for {
		println("Para sempre")
		time.Sleep(time.Second)
	}
}
