package main

import "fmt"

// CONSTANTE é sempre um valor fixo, não pode ser alterado.
const ebulicaoF float64 = 212.0

func main() {
	// VARIÁVEL é um valor que pode ser alterado durante a execução do programa.
	var tempF = ebulicaoF
	var tempC = (tempF - 32) * 5 / 9
	fmt.Println("Ponto de ebulição da água em Fahrenheit:", tempF)
	fmt.Println("Ponto de ebulição da água em Celsius:", tempC)
}
