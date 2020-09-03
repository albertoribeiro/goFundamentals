package main

import "fmt"

func main() {
	//var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[1234546576587] = "Maria"
	aprovados[5676890879867] = "Alberto"
	aprovados[4567576375567] = "Dayanne"
	aprovados[6735345342225] = "Zito"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[6735345342225])
	delete(aprovados, 4567576375567)
	fmt.Println(aprovados)
}
