package main

////---------- panic é uma função em Go que interrompe a execução normal do programa e inicia o processo de recuperação.
// Ela é usada para sinalizar um erro grave que não pode ser tratado de forma normal.
// Quando uma função entra em pânico, ela para imediatamente e começa a desfazer todas as chamadas de função pendentes, executando os
// comandos defer em ordem inversa.
// Isso é útil para liberar recursos ou fazer limpeza antes que o programa termine.
////---------- O pânico pode ser recuperado usando a função recover, que permite que o programa continue a execução normal após um pânico.

import "fmt"

func main() {
	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Pânico")
}
