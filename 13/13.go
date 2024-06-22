//Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 5
	b := 3
	fmt.Println("before change: a = ", a, ", ", "b = ", b)

	//Меняем местами значения.
	a, b = b, a
	fmt.Println("before change: a = ", a, ", ", "b = ", b)
}
