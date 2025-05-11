package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

func (r *retangulo) perimetro() int {
	return 2 * (r.comprimento + r.altura)
}

func main() {
	r := retangulo{comprimento: 10, altura: 5} // OU pode ser escrito assim: r := retangulo{10, 5}
	fmt.Println("A ÁREA = ", r.area())
	fmt.Println("O PERÍMETRO = ", r.perimetro())
}

// O método area() calcula a área do retângulo multiplicando o comprimento pela altura.
// O método perimetro() calcula o perímetro do retângulo somando o comprimento e a altura e multiplicando por 2.
// O método area() e o método perimetro() são definidos para o tipo retangulo.
// Eles podem ser chamados em qualquer instância do tipo retangulo.
// O método area() e o método perimetro() são definidos com um ponteiro para o tipo retangulo.
// Isso significa que eles podem modificar o valor do retângulo original.
// Se você não usar um ponteiro, o valor do retângulo original não será modificado.
// Isso é útil quando você deseja modificar o valor do retângulo original.
// Se você não usar um ponteiro, o valor do retângulo original não será modificado.
