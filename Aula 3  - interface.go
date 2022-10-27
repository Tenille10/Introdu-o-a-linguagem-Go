// Interfaces são coleções
// de métodos.

package main

import (
	"fmt"
	"math" //utilizar pq (pi)
)

// Aqui temos uma interface usada para formas geométricas.
type geometria interface {
	area() float64
	perimetro() float64
}

// Para nossa aula vamos usar essa interface nos
// tipos QUADRADO e CÍRCULO.
// área de um quadrado: lado² ou lado*lado
// perímetro = soma dos lados
type quadrado struct {
	lado float64
}
type círculo struct { //área círculo: (Pi)*raio² perímetro do círculo: 2*r*(pi)
	raio float64
}

// Para usar uma interface em Go, só precisamos
// usar todos os métodos da interface. Aqui nós
// usaremos `geometria` em `quadrado`s.
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// A implementação dos `círculo`s.
func (c círculo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c círculo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Se a variável tem um tipo interface, então nós podemos chamar
// métodos que estão na interface nomeada. Aqui temos uma
// função genérica `medida` tomando vantagem desse
// trabalho em qualquer `geometria`.
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado: 25}
	c := círculo{raio: 100}

	// Os tipos de estrutura `círculo` e `quadrado` ambos
	// implementam a interface `geometria` então nós podemos usar
	// instâncias dessas estruturas como argumentos para `medir`.

	medir(q)
	medir(c)
}
