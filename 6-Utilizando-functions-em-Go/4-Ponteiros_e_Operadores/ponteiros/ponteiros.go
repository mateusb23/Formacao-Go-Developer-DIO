package main

// Ponteiros são variáveis que armazenam o endereço de memória de outra variável.
// Eles são úteis quando você deseja modificar o valor de uma variável dentro de uma função sem precisar retornar o novo valor.
// Em Go, o operador & é usado para obter o endereço de memória de uma variável, enquanto o operador * é usado para
// desreferenciar um ponteiro e acessar o valor armazenado nesse endereço.
//
// Exemplo de uso de ponteiros em Go

import "fmt"

func inicial(xPtr *int) {
	*xPtr = 3
}

func main() {
	x := 5
	fmt.Println("Sem uso de ponteiros: ", x)
	inicial(&x)
	fmt.Println("Com uso de ponteiros: ", x)
	fmt.Println("Endereço de x:  ", &x)
}
