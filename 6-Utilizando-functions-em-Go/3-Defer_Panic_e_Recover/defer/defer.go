package main

//// defer é uma palavra-chave em Go que permite adiar a execução de uma função até que a função que a chamou retorne.
// defer escalona nossas funções

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda-feira")
}

func main() {
	defer dia1()
	dia2()
}
