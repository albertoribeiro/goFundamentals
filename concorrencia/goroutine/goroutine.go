package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	/*
		fale("Maria", "Pq vc não fala comigo?", 3)
		fale("João", "Só posso flar depois de vc!", 1)

		Maria: Pq vc não fala comigo? (iteração 1)
		Maria: Pq vc não fala comigo? (iteração 2)
		Maria: Pq vc não fala comigo? (iteração 3)
		João: Só posso flar depois de vc! (iteração 1)

		as funções foram executadas de foram serial.
	*/
	/*
		go fale("Maria", "Ei...", 50)
		go fale("João", "Opa...", 50)

		a thread so roda enquanto a thread principal estive ativa,
		no caso logo após a chamada com o go routine, acabou a thread main
		e as funções não foram executdas
		se colocarmos um sleep apos achamadas, as funções serão executadas até o fim do sleep

		time.sleep(time.second * 5)
	*/

	/*
		go fale("Maria", "Ei...", 50)
		fale("João", "Opa...", 50)

		neste caso, a go routine maria será executada até o fim da execução de fale"joao"
	*/
}
