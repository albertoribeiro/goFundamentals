package main

import "fmt"

func main() {
	r1, r2 := trocar(3, 4)
	fmt.Println(r1, r2)
}

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo, como ja foi definido o retorno nomeado ...
}
