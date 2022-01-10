package main

import (
	
	"fmt"
	"math/rand"
	"time"
)

//criar um pequeno bloco de inicialização de uma variável
// antes de executar o if
//pega nanossegundo da data atual do código

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	// s := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i> 5{
		fmt.Println("Ganhou!!!!")
	} else {
		fmt.Println("Perdeu!")
	}	
}