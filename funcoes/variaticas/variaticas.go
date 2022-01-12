package main

import "fmt"

/*funções variáticas: recebem quantidade de parâmetros variáveis
... indica função variatica e trabalhar dentro dos parametros com array
*/
func media(numeros ...float64) float64{
	total := 0.0
	for _, num := range numeros{
		total += num
	}
	//media dividio pelo tamanho do array(conversão do tamanho poque é inteiro)
	return total / float64(len(numeros))
}

func main(){
	fmt.Printf("Média: %.2f", media(7.7, 8.1, 5.9, 9.9))
}