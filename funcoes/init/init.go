package main

import "fmt"

//Init: função executada quando arquivo em GO é interpretado

func init(){
	fmt.Println("Inicializando...")
}

func main() {
	/* mesmo não sendo chamada, por convenção do Go
	o método init é chamado antes do main */
	fmt.Println("Main...")
}