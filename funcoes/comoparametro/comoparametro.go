package main

import "fmt"

func main() {
	resultado := exec(mult, 3, 4)
	fmt.Println(resultado)
}

func mult(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
