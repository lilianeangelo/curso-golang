package main

import (
	"fmt"
	"time"
)

//tipo interface promove trabalhar com dados mais genéricos
func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "desconhecido"	
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(3))
	fmt.Println(tipo("opa"))
	fmt.Println(tipo(func(){}))
	fmt.Println(tipo(time.Now()))
}