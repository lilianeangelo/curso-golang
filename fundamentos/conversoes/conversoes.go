package main

import (
	"fmt"
	c "strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x/ float64(y))

	nota := 6.9
	notaFinal:= int(nota)
	fmt.Println(notaFinal)

	//cuidado ao concatenar ... 97 é o código da tabela ASCII
	fmt.Println("Teste " + string(97))

	//int para string
	fmt.Println("Teste " + c.Itoa(123))

	/* string para int - em GO é possível retornar dois valores 
	(no caso abaixo um número e um erro) */
	num, _ := c.Atoi("123")
	fmt.Println(num - 122)

	b, _ := c.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}