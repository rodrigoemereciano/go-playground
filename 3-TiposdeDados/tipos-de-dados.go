package main

import "fmt"

func main() {

	// int8, int16, int32, int64

	var numero int64 = 100
	fmt.Println(numero)

	// uint // unsygned int
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// byte = uint8
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	texto := 5
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)

}
