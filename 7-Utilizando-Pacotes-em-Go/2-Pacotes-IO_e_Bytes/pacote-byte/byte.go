package main

import (
	"bytes"
	"fmt"
)

// Pacote Bytes
// O pacote Bytes é usado para trabalhar com bytes em Go.
// Ele fornece funções para manipular bytes, como comparar, copiar, etc.

func main() {
	fmt.Printf("%s", bytes.Title([]byte("Tudo bem contigo?")))
}
