// aceita um conjunto de dados e o reduz a um tamanho fixo menor.
// hashes são usados em TUDO em programação, principalmente para buscar DADOS e DETECTAR em Go dívidas ENCRIPTOGRAFADAS.
// listas das NÃO CRIP: adler32, crc32, crc64
// nesse código vamos criar uma HASH e escrever nossos dados nele.

package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// criando uma hash
	h := crc32.NewIEEE()
	// escrever nossos dados na hash
	h.Write([]byte("código com pacote hash"))
	// calcular o hash
	v := h.Sum32()
	fmt.Println(v)
}
