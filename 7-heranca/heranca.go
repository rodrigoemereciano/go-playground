package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Arquivo herança")

	p1 := pessoa{"João", "Silva", 30, 1.75}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "Unicamp"}

	fmt.Println(e1.nome)
	fmt.Println(e1.altura)
}
