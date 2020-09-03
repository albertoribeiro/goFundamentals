package main

import "fmt"

func main() {

	funcESalarios := map[string]float64{
		"José João":      11325.76,
		"Gabriela Silva": 15456.45,
		"Pedro Junior":   1200.76,
	}

	fmt.Println(funcESalarios)

	funcESalarios["Rafael Filho"] = 1350.0
	fmt.Println(funcESalarios)
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios { // o primeiro valor sempre será a chave do map
		fmt.Println(nome, salario)
	}
}
