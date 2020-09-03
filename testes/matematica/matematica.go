// Teste unitário - testar uma unidade, pode ser uma função,
// deve isolar a função, quanto mais isolada melhor

package matematica

import (
	"fmt"
	"strconv"
)

// Media calcula a média dos valores passados no parametro
func Media(numeros ...float64) float64 {
	total := 0.0
	for _, num := range numeros {
		total += num
	}
	media := total / float64(len(numeros))
	mediaArredondada, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", media), 64)
	return mediaArredondada
}
