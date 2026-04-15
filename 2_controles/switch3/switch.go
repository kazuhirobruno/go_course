package main

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "func"
	default:
		return "unknown"
	}
}

func main() {
	println(tipo(1))
	println(tipo(1.0))
	println(tipo("hello"))
	println(tipo(func() {}))
}
