package main

import "fmt"

func main() {
	// AULA sobre Loop infinito com Break e Continue
	i := 0
	for {
		if i < 20 {
			fmt.Println("i é menor que 20.")
			i++
		} else {
			break // O break interrompe o loop
		}
	}
}
