package main

import "fmt"


func multiplicacao(a, b int) int{
	return a * b
}
//função exec que recebe uma função que recebe dois parametros inteiros e retorna um inteiro
func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}
func main(){
	resultado := multiplicacao(3, 4)
	fmt.Println(resultado)
}