package main

import (
	"encoding/json"
	"fmt"
)

/**Convenção da linguagem Go:
Quando começa com minúscula é privado.
Quando começa com maíuscula é público**/
type produto struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
	Preco float64 `json:"preco"`
	Tags []string `json:"tags"`
}

func main(){
	//struct para json
	p1 := produto{1, "Geladeira", 2068.99, []string{"promocao", "eletrodomesticos", "Brastemp"}}
	//Marshal retorna dois valores: json como texto e segundo o erro
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	//json para struct
	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta","preco":7.56,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}