//Goroutines: é uma função capaz de executar de modo concorrente com outras funções
// palavra reservada para goroutine: go

package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
func main() {
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)

}
