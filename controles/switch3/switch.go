package main

import (
	"fmt"
	"time"
)

// pode ser util para tatar parametros genericos
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "interiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "funcao"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(9.3))
	fmt.Println(tipo(8))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
}
