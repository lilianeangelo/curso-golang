package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	/* switch: multiplas seleções- executou o case, ele sai do bloco switch
	 Em Java, é necessário escrever a palavra break para interromper o case
	 o fallthrough é para executar os outros debaixo -descer para os próximos*/
	switch nota{
	case 10:
		fallthrough 
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota Inválida!"						 
	}
}

func main(){
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.1))
	fmt.Println(notaParaConceito(1.1))
	fmt.Println(notaParaConceito(13.0))
}