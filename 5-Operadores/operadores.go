package main

import "fmt"

func main() {

	//Aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25

	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//Atribuicao

	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	//Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//Operadores Logicos
	fmt.Println(true && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Operadores Unarios
	numero := 10
	numero++
	numero += 15
	numero--
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	//Operadores Ternarios
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
