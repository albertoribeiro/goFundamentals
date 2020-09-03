package main

// Alias no import
import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 4.1415
	var raio = 3.2 // tipo flaot64 inferido pelo navegaros

	// forma reduzida de criar var
	// com : vc esta declarando e inicializando, criando variavel 
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)
	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
