package main

import "fmt"

// forçano o paramentro sem sinal, dscarta a necessidade de tratar valor negativo
func fatorial(n uint) uint {
	if n == 0 {
		return 1
	}
	fatorialAnterior := fatorial(n - 1)
	return n * fatorialAnterior
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
