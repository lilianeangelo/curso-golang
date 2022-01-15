package main

import "fmt"

func imprimirAprovados(aprovados ...string){
	fmt.Println("Lista de Aprovados!")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

/*funções variáticas não usam array porque como diz, são variaveis os parametros
logo só funciona com slice */

func main(){
	aprovados := []string{"Maria", "Pedro", "Augusto", "Valentina"}
	imprimirAprovados(aprovados...)
}