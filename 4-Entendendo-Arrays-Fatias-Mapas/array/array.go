package main

import "fmt"

func main() {
	var x [5]float64 // OU // x := [5]float64{} // Declaração de um array de inteiros com 5 elementos
	fmt.Println(x)   // Imprime o array vazio (inicializado com zeros)
	x[4] = 8.3       // Atribui o valor 10 ao índice 4 do array
	x[3] = 7.2
	x[1] = 5.4
	x[0] = 3.1
	fmt.Println(x) // Imprime o array após a atribuição

	total := 0.0
	for i := 0; i < len(x); i++ { // Loop para percorrer o array
		total += x[i] // Soma os elementos do array
	}
	fmt.Println("Valor da soma total dos elementos do Array = ", total)
	media := total / float64(len(x))                   // Calcula a média
	fmt.Println("A média de valor do Array = ", media) // Imprime a média dos elementos do array
}
