package main

func main() {
	i := 1

	var p *int = nil

	p = &i

	*p++

	i++

	println(*p, p, i, &i)
}
