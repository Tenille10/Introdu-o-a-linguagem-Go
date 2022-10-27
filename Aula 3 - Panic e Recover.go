//panic: erro do programados/erro execução tempo
//recover: ela interrompe o panic

package main

import "fmt"

func main() {
	defer func() {

		x := recover()
		fmt.Println(x)
	}()
	panic("Pânico")
}
