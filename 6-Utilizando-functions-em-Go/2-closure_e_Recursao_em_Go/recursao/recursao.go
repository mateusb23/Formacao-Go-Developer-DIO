/// Recrusão é uma técnica onde uma função chama a si mesma para resolver um problema.
// Isso é útil para resolver problemas que podem ser divididos em subproblemas menores, como calcular o
// fatorial de um número ou percorrer uma árvore.

package main

import "fmt"

func main() {
	fmt.Println(fatorial(5))
}

func fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}
