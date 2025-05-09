package main

import "fmt"

//// MAPAS EM GO
// Mapas são estruturas de dados que armazenam pares chave-valor, permitindo acesso rápido aos
// alores associados a uma chave específica.

// var x map[string]int // Declaração de um mapa vazio
// var x = make(map[string]int) // Criação de um mapa vazio
// var x = map[string]int{} // Criação de um mapa vazio
// var x = map[string]int{"chave1": 1, "chave2": 2} // Criação de um mapa com valores iniciais

func main() {

	x := make(map[string]int)
	x["chave1"] = 5 // Adiciona um par chave-valor ao mapa
	x["chave2"] = 8
	x["chave3"] = 35
	x["chave4"] = 6
	fmt.Println(x) // Imprime o mapa completo
	fmt.Println("-------------------------------------")
	fmt.Println("Mostre o valor associado à chave 3: ", x["chave3"])

	fmt.Println("-------------------------------------")

	y := make(map[int]int)
	y[1] = 9
	y[2] = 2
	y[3] = 7
	fmt.Println(y)
	fmt.Println("-------------------------------------")
	fmt.Println("Mostre o valor associado à chave 2: ", y[2])

}
