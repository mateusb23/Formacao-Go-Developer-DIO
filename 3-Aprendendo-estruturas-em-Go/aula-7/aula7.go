package main

import "fmt"

func main() {
	// Loop de Repetição HIERÁRQUICO
	for horas := 0; horas <= 12; horas++ {
		fmt.Println("Horas: ", horas)

		for min := 0; min < 60; min++ {
			fmt.Println(min)
		}

	}
}
