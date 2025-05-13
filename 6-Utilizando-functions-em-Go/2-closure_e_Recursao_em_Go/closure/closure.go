/// Go: Closure e Recursão
// Closure é uma função que captura o ambiente em que foi criada.
// Isso significa que ela pode acessar variáveis que estavam no escopo quando a função foi definida, mesmo depois que esse escopo já não existe mais.
// Isso é útil para criar funções que têm um estado interno, como contadores ou geradores de números aleatórios.
// Recursão é uma técnica onde uma função chama a si mesma para resolver um problema.

// Closure: é a capacidade de uma função "chamar" uma variável que está em outra função.

package main

import "fmt"

func main() {
	x := 0
	numero := func() int {
		x++
		return x
	}
	fmt.Println((numero()))
	fmt.Println((numero()))
	fmt.Println((numero()))
	fmt.Println((numero()))
	fmt.Println((numero()))
}
