//recursão é a capcidade de uma função chamar ela mesma
//calcular fatorial

package main

import "fmt"

func fatorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func main() {
	fmt.Println(fatorial(2))
}
