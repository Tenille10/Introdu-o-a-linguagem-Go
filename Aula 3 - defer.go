// defer: escalona nossas funções
package main

import "fmt"

func dia1() {
	fmt.Println("Domingo")
}

func dia2() {
	fmt.Println("Segunda-feira")
}

func main() {
	defer dia2()
	dia1()
}
