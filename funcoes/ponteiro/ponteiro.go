package main

import "fmt"
// Resumindo, passar ponteiro é a mesma coisa que passar parameto por referencia ao invés de passar por valor
func inc1(n int) {
	n++
}

// um ponteiro é representado por *
func inc2(n *int) {
	// * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n)
	fmt.Println(n)

	// & é usado para pegar o endereço da variavel
	inc2(&n) // por referencia
	fmt.Println(n)
}
