package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 10
	array1[1] = 20
	array1[2] = 30
	array1[3] = 40
	array1[4] = 50
	fmt.Println(array1)

	array2 := [5]string{"Go", "Java", "Python", "JavaScript", "C#"}
	fmt.Println(array2)

	array3 := [...]float32{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(array3)

	// com o slice nao precisamos definir o tamanho
	slice := []int{100, 200, 300, 400, 500, 600}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 700)
	fmt.Println(slice)

	slice2 := array2[:3]
	fmt.Println(slice2)

	array2[1] = "TypeScript"
	fmt.Println(slice2)
}
