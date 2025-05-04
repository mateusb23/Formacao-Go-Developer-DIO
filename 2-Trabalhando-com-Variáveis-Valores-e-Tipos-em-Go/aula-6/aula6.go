package main

import "fmt"

var a int = 20
var b string = "Nome"

func main() {
	// para converter um tipo para outro, é necessário fazer a conversão explícita
	// ou seja, não é possível fazer a conversão implícita
	// var c float64 = a  // não é possível fazer a conversão dessa forma no Go, pois daria erro
	// ---> FORMA CORRETA DE CONVERSÃO DE TIPO <---
	var c float64 = float64(a) // conversão de int para float64
	fmt.Printf("O valor de c = %g \nO valor de b = %s", c, b)
}
