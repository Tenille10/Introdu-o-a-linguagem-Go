package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("AulaGo.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
