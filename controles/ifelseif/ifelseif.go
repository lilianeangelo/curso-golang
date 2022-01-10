package main

import "fmt"

func notaParaConceito(n float64) string{
	if n > 10{
		return "Não existem notas acima de 10 para conceitos. Digite novamente!"	
	} else if n >= 9 && n <= 10 {
		return "A"
	} else if n >= 8 && n < 9 {
		return "B"
	} else if n >= 5 && n < 8{
		return "C"
	} else if n >= 3 && n < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(12.0))
	fmt.Println(notaParaConceito(9.0))
	fmt.Println(notaParaConceito(6.0))
	fmt.Println(notaParaConceito(8.0))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(2.0))
}