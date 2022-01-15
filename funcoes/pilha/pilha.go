package main

import "runtime/debug"

func f3(){
	//FIFO: First In, First Out
	//imprimir a pilha de execução de um programa no momento que a função foi chamada
	debug.PrintStack()
}

func f2(){
	f3()
}

func f1(){
	f2()
}

func main(){
	f1()
}