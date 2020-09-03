package main

import "fmt"

// funções com quantidade de parametros variado
func main() {
	fmt.Printf("Média: %.2f", media(7.7, 2.3, 6.5, 8.4, 9.4))
}
func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}
	return total / float64(len(numeros))
}
