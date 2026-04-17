package main

import "math"

// Reta é definida por dois pontos
type Ponto struct {
	x float64
	y float64
}

func Catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

func Distancia(a, b Ponto) float64 {
	cx, cy := Catetos(a, b)
	return math.Sqrt(cx*cx + cy*cy)
}
