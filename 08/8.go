//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func SetIBit(n int64, i int, val int) int64 {
	//Для смены значения на i-ой позиции на 1-цу используется двоичный оператор ИЛИ (|=).
	if val == 1 {
		n |= 1 << i
	} else {
		//Для того, чтобы преобразовать любое исходное значение в 0, используется И-НЕТ (&^): конъюнкция с 0 всегда дает 0.
		n &^= 1 << i
	}
	return n
}

func main() {
	var num int64
	var i, val int

	fmt.Println("Consequtevely write number, bit's position and value you want to set")
	fmt.Scan(&num, &i, &val)

	if i < 0 || i > 63 {
		fmt.Println("Wrong position of bit")
		return
	}

	if val != 0 && val != 1 {
		fmt.Println("Wrong value of a bit (has to be 1 or 0)")
		return
	}

	fmt.Println(SetIBit(num, i, val))
}
