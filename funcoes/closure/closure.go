package main

import "fmt"

/*closure: fechamento
função que envolve outra função
função interna entende aonde foi declarada
*/
func closure() func(){
	x := 10
	var funcao = func ()  {
		fmt.Println(x)
	}
	return funcao
}

func main(){
	x := 20
	fmt.Println(x)

	//ler a função nas origens que foi feita
	imprimeX := closure()
	imprimeX()
}
//stackoverflow: estouro de pilha