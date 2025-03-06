package main

import "fmt"

func somar(a int8, b int8) int8 {
	return a + b
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
}
