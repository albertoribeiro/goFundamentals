package main

import "fmt"

// funções com quantidade de parametros variado
func main() {
	aprovados := []string{"Maria", "João", "Francisco", "Renata"}
	imprimirAprovados(aprovados...)
}
func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}
