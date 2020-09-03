package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1

	for i <= 10 { // não existe while em GO
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Println(i)
	}

	fmt.Println("\nMisturando")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	// laço infiinito
	for {
		fmt.Println("Para Sempre...")
		time.Sleep(time.Second * 5)
	}

}
