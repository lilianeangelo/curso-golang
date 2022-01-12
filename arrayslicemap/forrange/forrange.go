package main

import "fmt"

//for com range
func main(){
	numeros := [...]int{1, 2, 3, 4, 5} //compilador conta

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	//acessar o número sem precisar usar o índice
	for _, num := range numeros{
		fmt.Println(num)
	}
}