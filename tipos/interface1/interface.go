package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct{
	nome string
	sobrenome string
}

type produto struct {
	nome string
	preco float64
}

//Polimorfismo mesmo Go não sendo Orientada à objeto
// interfaces são implementadas implicitamente
func (p pessoa) toString() string {
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string {
	//Sprintf para armazenar como uma string porque o retorno é uma string
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}

func imprimir(x imprimivel){
	fmt.Println(x.toString())
}

func main(){
	var coisa imprimivel = pessoa{"Cassia", "Jane"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Saia", 69.99}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	p2 := produto{"Saia", 169.99}
	imprimir(p2)
}