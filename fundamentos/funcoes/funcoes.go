package main

import "fmt"

func main() {
	resultado := somar(43, 56)
	imprimir(resultado)
}

func somar(a int, b int) int {
	return a + b
}

func imprimir(valor int) {
	fmt.Println(valor)
}
