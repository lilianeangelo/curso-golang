package main

import (
	"fmt"
	"time"
)

//unica estrutura de repetição do Go
//não tem while em Go

func main() {
	i := 1

	for i <= 10 {
		fmt.Println(i)
		i++
	}
	
	for i := 0; i <= 20; i +=2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\nMisturando... ")
	for i := 1; i <= 10; i++{
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	//laço infinito
	for{
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}