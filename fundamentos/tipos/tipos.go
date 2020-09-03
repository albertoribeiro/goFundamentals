package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é ", reflect.TypeOf(32000))

	// sem sinal (só positivo)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("o byte é ", reflect.TypeOf(b))

	// com sinal int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("o valor maximo de i1 é ", i1)
	fmt.Println("o tipo de i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("o rune de a é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("o tipo de x é ", reflect.TypeOf(x))
	fmt.Println("o tipo de literal de 49.99 é  ", reflect.TypeOf(49.99)) // float64

	// Boolean
	bo := true
	fmt.Println("o tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá meu nome é Alberto "
	fmt.Println(s1 + "!")
	fmt.Println("o tamanho da strng é", len(s1))

	s2 := `Olá  
	meu 
	nome 
	é 
	Alberto`
	fmt.Println("o tamanho da strng é", len(s2))

	// não possui o tipo char.

}
