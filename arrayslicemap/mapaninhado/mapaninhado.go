package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Machado": 54216.30,
			"Galvão Bueno": 60415.30,
		},
		"M":{
			"Mauricio Mendes": 22351.48,
			"Monica Veloso": 30504.91,
		},
		"P": {
			"Pablo Ventura": 8541.75,
			"Patricia Pillar": 9654.60,
		},
	}
	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra{
		fmt.Println(letra, funcs)
	}
}