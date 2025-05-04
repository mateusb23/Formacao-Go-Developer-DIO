package main

import "fmt"

func main() {
	// Tamanho da string
	// A string é um slice de bytes, então o tamanho dela é o número de bytes
	fmt.Println(len("Hello, World!"))
	// Acessando o terceiro caractere da string
	fmt.Println("Hello, World!"[2])
	// Somando duas strings
	// A soma de duas strings é a concatenação delas
	fmt.Println("Hello" + "World!")
}
