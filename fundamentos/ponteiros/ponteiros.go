package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço de memoria da variavel i
	*p++   // desreferenciando (pegando o valor)
	i++

	// Go não tem aritimetica de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)

}
