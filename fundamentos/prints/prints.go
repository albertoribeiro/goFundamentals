package main

import "fmt"

func main() {
	// Print - imprime e permanece na mesma linha
	fmt.Print("Mesma ")
	fmt.Print("linha!")

	// Print - imprime e quebra a  linha
	fmt.Println("Nova ")
	fmt.Println("linha!")

	x := 3.141516

	// fmt.Println("O valor de x é" + x) - erro de conversão
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)

}
