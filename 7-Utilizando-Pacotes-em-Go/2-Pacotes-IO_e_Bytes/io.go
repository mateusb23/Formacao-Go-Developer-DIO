package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Boa noite!"); err != nil {
		log.Fatal(err)
	}
}
