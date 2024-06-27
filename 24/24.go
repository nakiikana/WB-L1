// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Искапсуляция в рамках одного пакета реализуется за счет переменной, написанной со строчной буквы.
type Point struct {
	x float64
	y float64
}

// Конструктор структуры точки.
func newPoint(x float64, y float64) Point {
	return Point{x: x, y: y}
}

func (p1 *Point) FindDistance(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := newPoint(1.0, 2.0)
	p2 := newPoint(3.0, 9.0)
	fmt.Printf("p1: %v, p2: %v - dist: %v\n", p1, p2, p2.FindDistance(p1))
}
