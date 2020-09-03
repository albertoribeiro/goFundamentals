package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)        // cria o array intrno com 20 elementos
	fmt.Println(s, len(s), cap(s)) // cap = capacidade

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s)) // cap = capacidade

	s = append(s, 1)
	fmt.Println(s, len(s), cap(s)) //o slice cresce automaticamente
}
