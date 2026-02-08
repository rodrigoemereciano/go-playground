package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":     "Rômulo",
		"endereço": "Rua 01",
	}

	fmt.Println(usuario)

	usuario["telefone"] = "(11) 99999-9999"
	fmt.Println(usuario)

	usuario["idade"] = "25"
	fmt.Println(usuario)

	delete(usuario, "sobrenome")
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Rômulo",
		},
		"endereço": {
			"logradouro": "Rua 01",
			"numero":     "0",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "endereço")
	fmt.Println(usuario2)

}
