package main

import (
	"fmt"

	"github.com/kazuhirobruno/area/area"
)

func main() {
	fmt.Println(area.Circle(5.0))    // Deve imprimir 78.53981633974483
	fmt.Println(area.Rect(4.0, 6.0)) // Deve imprimir 24.0
}
