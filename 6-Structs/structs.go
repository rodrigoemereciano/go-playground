package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Rômulo"
	u.idade = 25

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	usuario2 := usuario{"Davi", 23, enderecoExemplo}
	fmt.Println(usuario2)
}
