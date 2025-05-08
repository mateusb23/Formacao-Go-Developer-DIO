package main

// MOSTRANDO DIFERENÇAS ENTRE ARRAYS E FATIAS (slices) EM GO
// Arrays são de tamanho fixo, enquanto fatias (slices) podem crescer e encolher dinamicamente.
// var x []float64 // Declaração de uma fatia vazia
// fatia := make([]float64, 4) // Criação de uma fatia com 4 elementos
// fatia := [low:high] // Criação de uma fatia a partir de um array
// fatia := arr[0:5]

import "fmt"

func main() {
	// Declaração de uma fatia vazia
	var x []float64 // Declaração de uma fatia vazia
	fmt.Println(x)  // Imprime a fatia vazia (nil)

	// Criação de uma fatia com 4 elementos
	fatia := make([]float64, 5) // Criação de uma fatia com 4 elementos
	fmt.Println(fatia)          // Imprime a fatia inicializada com zeros

	// Atribuição de valores à fatia
	fatia[0] = 1.1
	fatia[1] = 2.2
	fatia[2] = 3.3
	fatia[3] = 4.4
	fatia[4] = 5.5

	// Imprime a fatia após a atribuição
	fmt.Println(fatia) // Imprime a fatia após a atribuição

	// Criação de uma fatia a partir de um array
	arr := [5]float64{1.0, 2.0, 3.0, 4.0, 5.0} // Criação de um array com 5 elementos do tipo float64
	fatia2 := arr[0:3]                         // Criação de uma fatia a partir do array (elementos 0 a 2)
	fmt.Println(fatia2)                        // Imprime a fatia criada a partir do array

	// Exemplo de fatiamento com low e high
	fatia3 := arr[1:4]  // Criação de uma nova fatia (elementos 1 a 3)
	fmt.Println(fatia3) // Imprime a nova fatia criada

	// Exemplo de fatiamento com low e high (sem especificar o high)
	fatia4 := arr[:3]   // Criação de uma nova fatia (elementos 0 a 2)
	fmt.Println(fatia4) // Imprime a nova fatia criada

	// Exemplo de fatiamento com low e high (sem especificar o low)
	fatia5 := arr[2:]   // Criação de uma nova fatia (elementos 2 até o final)
	fmt.Println(fatia5) // Imprime a nova fatia criada

	// Exemplo de fatiamento com low e high (sem especificar o low e o high)
	fatia6 := arr[:]    // Criação de uma nova fatia
	fmt.Println(fatia6) // Imprime a nova fatia criada

}

// ### Resumo:
// - Toda fatia é baseada em um array subjacente, mas nem toda fatia é um array.
// - Arrays são fixos e independentes, enquanto fatias são flexíveis e dependem de arrays subjacentes.
// A fatia é uma abstração de um array, que permite trabalhar com arrays de forma mais flexível e dinâmica.
// As fatias são mais comuns em Go do que os arrays, pois oferecem mais funcionalidades e são mais fáceis de usar.
// Além disso, as fatias podem ser passadas como argumentos para funções, o que facilita a manipulação de dados em Go.
