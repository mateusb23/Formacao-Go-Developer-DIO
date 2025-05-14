//// Pacotes (Caixinha de funções)
// Pacotes são como caixinhas de funções que podemos usar em nossos programas.

/// ---> função contains
// contains verifica se uma substring está contida em uma string.
// exemplo: "abc" está contido em "abcdefg"

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("abcdefg", "abc")) // true
	fmt.Println(strings.Contains("abcdefg", "xyz")) // false
}
