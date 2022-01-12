package main

import "fmt"

/*Defer atrasa a execução de 
alguma parte do código*/
func obterValorAprovado(numero int) int {
	defer fmt.Println("fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	}
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obterValorAprovado(10000))
	fmt.Println(obterValorAprovado(2000))
}