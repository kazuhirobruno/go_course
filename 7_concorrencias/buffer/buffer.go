package main

import "fmt"

func rotina(c chan int) {
	fmt.Println("Executou")
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	c <- 6
}

func main() {
	ch := make(chan int, 3)
	ch2 := make(chan int, 5)
	go rotina(ch)
	go rotina(ch2)

	fmt.Println(<-ch)
	fmt.Println(<-ch2)
}
