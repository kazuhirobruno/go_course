package main

import "encoding/json"

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{
		ID:    1,
		Nome:  "Notebook",
		Preco: 2500.00,
		Tags:  []string{"eletrônicos", "computadores"},
	}
	p1Json, _ := json.Marshal(p1)
	println(string(p1Json))

	var p2 produto
	jsonString := `{"id":2,"nome":"Smartphone","preco":1500.00,"tags":["eletrônicos","celulares"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	println(p2.Nome)
}
