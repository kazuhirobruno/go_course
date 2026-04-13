package main

func comprar(trabalho1 bool, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2
	comprarSorvete := trabalho1 || trabalho2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	println("Tv 50\"?", tv50)
	println("Tv 32\"?", tv32)
	println("Sorvete?", sorvete)
}
