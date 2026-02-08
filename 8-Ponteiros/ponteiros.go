package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var2++
	fmt.Println(var1, var2)

	//PONTEIRO É UMA VARIÁVEL QUE ARMAZENA O ENDEREÇO DE MEMÓRIA DE OUTRA VARIÁVEL
	var var3 int = 10
	var ponteiro *int = &var3

	fmt.Println(var3, ponteiro)

	fmt.Println(var3, *ponteiro) // desreferenciando o ponteiro
}
