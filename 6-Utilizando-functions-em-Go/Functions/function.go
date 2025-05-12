// Função também é chamada de método
// Função também é chamado de procedimento ou sub-rotina
// Ela pega um valor de entrada, processa e retorna um valor de saída.
// Funções são declaradas com a palavra-chave func
// Funções podem ter parâmetros e retornar valores

package main

import "fmt"

func main() { // programa que calcula a média de uma sala
	lista_de_notas := []float64{98, 93, 77, 82, 83} // lista de notas
	media := media(lista_de_notas)                  // chamada da função media
	fmt.Println("A média da sala é:", media)        // imprime a média
}

func media(notas []float64) float64 { // função que calcula a média
	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	//// OU com for range ------
	// for _, nota := range notas {
	// 	total += nota
	// }
	return total / float64(len(notas)) // retorna a média
}
