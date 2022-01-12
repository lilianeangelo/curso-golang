package main

import "fmt"

/* os arrays são estruturas homogêneas
(mesmo tipo de dados dentro delas)
e uma estrutura estática (fixa: se o array tem
10 posições, ele continuará com 10 posições,
ou seja, não cresce seu tamanho)
*/
func main(){
	var notas [3]float64
	fmt.Println(notas)

	//array acessa por índices
	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	total := 0.0
	for i := 0; i < len(notas); i++{
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}