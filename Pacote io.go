package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if _, err := io.WriteString(os.Stdout, "Bom dia"); err != nil {
		log.Fatal(err)
	}
}
