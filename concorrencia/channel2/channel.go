package main

import (
	"fmt"
	"time"
)

// Channel é a forma de comunicação entre as goroutines.
// ponto de sincronismo de varios processos independentes.
// é um tipo

func doisTresQuatroVees(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go doisTresQuatroVees(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println("B")

	fmt.Println(<-c)

	// se eu tentar ler mais uma vez do canal , vai ocorrer deadlock.

}
