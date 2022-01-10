package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmética de ponteiros

	//endereço de memória = & 
	// nil é null
	// ponteiro aponta para outro dado de memória

	var p *int = nil

	// p guarda endereço de memória - pegando o endereço da variável
	p = &i
	*p++ // desreferenciando (pegando o valor)
	i++

	//o endereço é o compilador quem decide
	fmt.Println(p, *p, i, &i)


}