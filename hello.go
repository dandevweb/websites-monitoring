package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Danilo"
	version := 1.1
	age := 25

	fmt.Println("Hello,", name, ", sua idade é", age, "anos")
	fmt.Println("Este programa está na versão,", version)

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(name))
	fmt.Println("O tipo da variável version é", reflect.TypeOf(version))
	fmt.Println("O tipo da variável age é", reflect.TypeOf(age))
}
