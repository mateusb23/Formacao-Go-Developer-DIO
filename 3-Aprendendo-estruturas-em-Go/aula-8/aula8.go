package main

import "fmt"

func main() {
	// Loop em que o for funciona como um while
	// O for é o único loop em Go, não existe while ou do while
	i := 0
	for i < 20 {
		fmt.Println("i é menor que 20.")
		i++
	}
}
