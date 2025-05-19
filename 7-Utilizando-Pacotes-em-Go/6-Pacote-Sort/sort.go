// Ordena dados de forma crescente e decrescente.
// O pacote sort implementa a interface sort.Interface, que é usada para ordenar dados.
// A interface sort.Interface tem três métodos: Len(), Less(i, j int) bool e Swap(i, j int).
// O método Len() retorna o número de elementos a serem ordenados.
// O método Less(i, j int) bool retorna verdadeiro se o elemento i deve vir antes do elemento j.
// O método Swap(i, j int) troca os elementos i e j.
// O pacote sort também fornece funções para ordenar slices e arrays de tipos primitivos, como int, float64 e string.

package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}

func (ps ParaNome) Less(i, j int) bool { // verifica se o nome do elemento na posição i é menor que o nome do elemento na posição j
	return ps[i].Nome < ps[j].Nome
}

func (ps ParaNome) Swap(i, j int) { // troca os elementos na posição i e j
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	criancas := []Dados{
		{"Julia", 5},
		{"João", 10},
	}
	sort.Sort(ParaNome(criancas)) // ordena os dados
	fmt.Println(criancas)         // imprime os dados ordenados
}
