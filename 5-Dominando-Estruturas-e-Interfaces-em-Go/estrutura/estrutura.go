package main

// Estrutura é um tipo de dado que agrupa diferentes tipos de dados
// em uma única unidade. É semelhante a uma classe em outras linguagens de programação.

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	fmt.Println(pessoa{"Ana", 54})
	fmt.Println(pessoa{"João", 20})
}
