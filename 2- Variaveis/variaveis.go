package main

import "fmt"

func main(){

	var variavel1 string = "Variável 1"
	var variavel2 string = "Variável 2"
	var variavel3 string = "Variável 3"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)

	var variavel1, variavel2, variavel3 string = "Variável 1", "Variável 2", "Variável 3"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)

	variavel1, variavel2, variavel3 := "Variável 1", "Variável 2", "Variável 3"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)

	var (
		variavel1 string = "Variável 1"
		variavel2 string = "Variável 2"
		variavel3 string = "Variável 3"
	)
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)

	variavel1, variavel2, variavel3 := "Variável 1", "Variável 2", "Variável 3"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)

}