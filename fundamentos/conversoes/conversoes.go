package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// fmt.Println(x / y) -- não permite operação entre tipos diferentes
	// é preciso converter explicitamente
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// CUIDADO
	// string(123) não converte o valor inteiro para string, converte para o caracter na tabela asc, no caso  "{"
	fmt.Println("Teste " + string(123))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))
	fmt.Println("Teste " + fmt.Sprint(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string para boolean
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}

	c, _ := strconv.ParseBool("1")
	if c {
		fmt.Println("Verdadeiro")
	}

	d, _ := strconv.ParseBool("0")
	if !d {
		fmt.Println("Falso")
	}

}
