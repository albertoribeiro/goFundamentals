package main

import "fmt"

func main() {

	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15456.00,
			"Guga Pereira":   23453.87,
		},
		"J": {
			"José João": 11324.43,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}
	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")
	fmt.Println(funcsPorLetra)

	for letra, funcs := range funcsPorLetra { // o primeiro valor sempre será a chave do map
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}

	}

}
