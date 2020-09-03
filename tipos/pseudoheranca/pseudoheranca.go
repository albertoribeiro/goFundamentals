package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro       // campos anonimos, pseudo herança, na verdade é composição
	turboligado bool
}

func main() {
	f := ferrari{}
	f.nome = "F40"
	f.velocidadeAtual = 0
	f.turboligado = true
	fmt.Printf("A ferrari %s está com o turbo ligado? %v\n", f.nome, f.turboligado)
	// através da composição tem acesso de leitura e escrita aos campos da struct carro
}
