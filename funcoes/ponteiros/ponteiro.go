package main

import "fmt"

//Ponteiros é representado por um *

func inc1(n int){
	n++
}

func inc2(n *int) {
	/*asterisco (*) é para desreferenciar para ter 
	acesso ao valor que o ponteiro referencia */
	*n++
}

func main(){
	n := 1

	inc1(n)
	fmt.Println(n)

	// & é usado para obter endereço da variável
	inc2(&n)
	fmt.Println(n)
}