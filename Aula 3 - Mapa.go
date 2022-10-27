package main

import "fmt"

func main() {

	//	x := make(map[string]int)
	//	x["chave"] = 10
	//	fmt.Println(x["chave"])
	//x := make(map[int]int)
	//x[1] = 20
	//x[2] = 30
	//fmt.Println(x[1], x[2])

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	fmt.Println(elemento["Li"])
}

//coleção ordenada (lista) pares chave-valor
//var x map[string]int
//x é um mapa de strings para ints
