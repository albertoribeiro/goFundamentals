package main

import "math"

// Iniciando com letra maiúscula é PÚBLICO (visibilidade dentro e fora do pacote) !
// Iniciando com letra minuscula é PRIVADO (visivel apenas dentro do pacote)!
// lembrando que pacote é diferente de arquivo

// Por exemplo...
// Ponto - gerará algo público
// ponto ou _Ponto - gerará algo privado

// Ponto  no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)
	return
}

// Distancia entre pontos
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

/*
Quando for necessário executar múltiplos arquivos no Windows, a forma correta é:
go run arquivo1.go arquivo2.go
*/
