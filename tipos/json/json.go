package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {

	// struct -> json
	p1 := produto{1, "Notebook", 1899.94, []string{"Promoção", "Eletrônico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json -> struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":14.94,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[0])

}

/*
Quando for necessário executar múltiplos arquivos no Windows, a forma correta é:
go run arquivo1.go arquivo2.go
*/
