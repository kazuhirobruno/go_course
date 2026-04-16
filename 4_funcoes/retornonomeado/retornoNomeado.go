package main

func troca(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return
}

func main() {
	p1, p2 := troca(10, 20)
	println(p1, p2)
	segundo, primeiro := troca(30, 40)
	println(segundo, primeiro)
}
