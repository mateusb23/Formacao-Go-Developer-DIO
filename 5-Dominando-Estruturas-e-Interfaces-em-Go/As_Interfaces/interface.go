package main

import (
	"fmt"
	"math"
)

// Aqui temos uma interface chamada geometria que define dois métodos: area() e perimetro().
// Esses métodos não têm implementação, apenas definem o que a interface deve ter.
// A interface é um tipo de dado que define um conjunto de métodos.
type geometria interface {
	area() float64
	perimetro() float64
}

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

func (q *quadrado) area() float64 {
	return q.lado * q.lado
}

func (q *quadrado) perimetro() float64 {
	return 4 * q.lado
}

func (c *circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

func (c *circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado: 10}
	c := circulo{raio: 5}

	medir(&q)
	medir(&c)
}
