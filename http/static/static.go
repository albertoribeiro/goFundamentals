package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}

// tem que rodar pela linha de comando, NÂO FUNCIONA PELO CODE RUNNER
