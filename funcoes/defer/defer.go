package main

import "fmt"

// adia a execução do comando para o termino do método
func obterValorAprovado(num int) int {
	defer fmt.Println("Fim") // executa este comando antes do return, após os outros comandos dentro do metodo
	if num > 5000 {
		fmt.Println("Valor Alto...")
		return 500
	}
	fmt.Println("Valor Baixo...")
	return num
}

func main() {
	fmt.Println(obterValorAprovado(5001))
	fmt.Println(obterValorAprovado(456))
}
