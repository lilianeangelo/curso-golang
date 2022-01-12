package main

import "fmt"

func main() {
	//todos os elemtentos precisam da vírgula
	funcsESalarios := map[string]float64{
		"José Augusto": 11562.30,
		"Vanessa Souza": 25684.60,
		"Antonio de Paula": 5482.15,
	}

	funcsESalarios["Leticia Vilela"] = 2140.16
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios{
		fmt.Println(nome, salario)
	}
}