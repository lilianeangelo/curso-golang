package main

import "fmt"

func main() {
	
	/* map: chave e valor - armazenando dentro do código 
	mapas devem ser inicializados - precisa ser
	%s: string; %d: inteiro; %f: float
	_ para ignorar um dos parâmetros
	*/ 

	aprovados := make(map[int]string)

	aprovados[12345678948] = "Maria"
	aprovados[43456897512] = "João"
	aprovados[84787569213] = "Leticia"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[84787569213])
	delete(aprovados, 84787569213)
	fmt.Println("Deleção abaixo...")
	fmt.Println(aprovados[84787569213])
}