package main

import "fmt"

/*o metodo make permite inicialização de um slice
com tamanho definido de pré elementos e capacidade
*/
func main(){
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	//len: tamanho do array, cap; capacidade (tamanho do array interno)
	fmt.Println(s, len(s), cap(s))
	//anexar ao slice mais 10 elementos
	s = append(s, 1, 2, 3, 4, 5, 6, 7 , 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	/*o slice não retorna um erro,
	ele internamente vai referenciando (expandindo)
	se você inserir um número a mais que a capacidade
	(dobra o tamanha da capacidade)*/

	s = append(s, 23)
	fmt.Println(s, len(s), cap(s))
}