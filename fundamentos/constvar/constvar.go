package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida para criar uma var
	// variável área não existe ainda (:=) estou declarando e inicializando a variável
	// uma vez que define uma variável, você precisa usar (no Go gera erro de compilação)
	//Pow = potenciação 
	area := PI * m.Pow(raio, 2)
	fmt.Println("A área da circunferência é:", area)

	//é possível criar blocos de constantes ou variáveis
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	//declaração de várias variáveis na mesma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"
	fmt.Println(g,h,i)
}