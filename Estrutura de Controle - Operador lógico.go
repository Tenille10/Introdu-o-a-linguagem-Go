package main

import "fmt"

func main() {

	//	x := 4

	//	if x == 2 || x == 3 {
	//		fmt.Println("Sim, x é 2 OU 3!")
	//	} else {
	//		fmt.Println("X não é nem 2 nem 3!")

	x := 9

	if x%2 == 0 && x%3 == 0 {
		fmt.Println("X é multiplo de 2 e 3")
	} else {
		fmt.Println("X não multiplo de 2 nem 3!")

	}
}
