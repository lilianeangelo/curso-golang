package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	abrirTetoSolar()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
	//pode adicionar mais metodos que queira
	apresentarNaFeira()
}

type bmw7 struct{}

func (b bmw7) ligarTurbo(){
	fmt.Println("Turbo...")
}

func (b bmw7) abrirTetoSolar(){
	fmt.Println("Teto Solar...")
}

func (b bmw7) apresentarNaFeira(){
	fmt.Println("Salão do automóvel...")
}

func main(){
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.abrirTetoSolar()
	b.apresentarNaFeira()
}