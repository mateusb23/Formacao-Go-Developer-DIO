// implementa rotinas para manipular arquivos e diretórios
// e para manipular caminhos de arquivos

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error { // a função percorre o diretório atual e mostra os arquivos que contém o "."
		fmt.Println(path)
		return nil
	})
}
