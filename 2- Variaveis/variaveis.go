package main

import "fmt"

func main() {
	var var1 string = "Variável 1"
	var var2 string = "Variável 2"
	var var3 string = "Variável 3"
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

	var var13, var14, var15 string = "Variável 13", "Variável 14", "Variável 15"
	fmt.Println(var13)
	fmt.Println(var14)
	fmt.Println(var15)

	var4, var5, var6 := "Variável 4", "Variável 5", "Variável 6"
	fmt.Println(var4)
	fmt.Println(var5)
	fmt.Println(var6)

	var (
		var7 string = "Variável 7"
		var8 string = "Variável 8"
		var9 string = "Variável 9"
	)
	fmt.Println(var7)
	fmt.Println(var8)
	fmt.Println(var9)

	var10, var11, var12 := "Variável 10", "Variável 11", "Variável 12"
	fmt.Println(var10)
	fmt.Println(var11)
	fmt.Println(var12)

	const const1 string = "Constante 1"
	const const2 string = "Constante 2"
	const const3 string = "Constante 3"
	fmt.Println(const1)
	fmt.Println(const2)
	fmt.Println(const3)
}
