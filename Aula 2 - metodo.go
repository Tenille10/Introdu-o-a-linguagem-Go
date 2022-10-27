// método é uma função associada a um tipo particular.
// Isto é, em POO(Prog. Orientada a Objeto), objeto é um valor(variável) e o método
// é uma função associada a esse objeto
package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

// Este método `área` possui um _tipo `retangulo`.
func (r *retangulo) area() int {
	return r.comprimento * r.altura
}

// Métodos podem ser definidos por qualquer tipo de receptor
// ponteiro ou valor. Aqui temos um exemplo do receptor de valor.
func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 50, altura: 25}

	// Aqui chamamos os 2 métodos definidos para a nossa estrutura.
	fmt.Println("Área: ", r.area())
	fmt.Println("Perímetro:", r.perimetro())

}
