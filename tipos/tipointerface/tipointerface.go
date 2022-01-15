package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println(coisa)

	coisa = 6
	fmt.Println(coisa)

	//interface pode receber vários tipos - bem genéricos como Object do Java
	type dinamico interface{}

	var coisa2 dinamico = "Ola"
	fmt.Println(coisa2)

	coisa2 = true
	fmt.Println(coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println(coisa2)
}