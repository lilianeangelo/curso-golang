package main

import (
	"fmt"
	"math"
)

func main() {

	a := 3
	b := 2
	fmt.Println("Soma =>", a+b)
	fmt.Println("Subtração =>", a-b)
	fmt.Println("Divisão =>", a/b)
	fmt.Println("Multiplicação =>", a*b)
	fmt.Println("Módulo =>", a%b) // resto da divisão

	//bitwise - bit a bit - valores binários
	fmt.Println("E =>", a&b)   //11 & 10 = 10 -- valor final da operação
	fmt.Println("OU =>", a|b)  //11 | 10 = 11
	fmt.Println("XOR =>", a^b) //11 ^ 10 = 01

	c := 3.0
	d := 2.0

	//outras operações usando math
	fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
	fmt.Println("Menor =>", math.Min(c, d))
	fmt.Println("Exponenciação =>", math.Pow(c, d))
}
