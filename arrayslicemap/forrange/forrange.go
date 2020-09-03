package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // o compilador que conta o tamanho do array

	// função range retorna dois valores, o index e o valor contido no index
	// semelhante ao foreach
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}

	for _, numero := range numeros {
		fmt.Printf(" %d\n", numero)
	}
}
