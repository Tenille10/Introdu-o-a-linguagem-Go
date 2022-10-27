package main

import (
	"io/ioutil"
	"log"
)

func main() {
	message := []byte("Teste go!")
	err := ioutil.WriteFile("teste", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
