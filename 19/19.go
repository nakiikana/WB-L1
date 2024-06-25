//Разработать программу, которая переворачивает подаваемую на ход
//строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func ReverseString(str string) string {
	//Переводим полученную строку к массиву рун.
	runes := []rune(str)
	mid := int(len(runes) / 2)
	//Меняем местами элементы, расположенные симметрично относительно середины массива.
	for i, j := 0, len(runes)-1; i <= mid && j >= mid; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(ReverseString("cat"))
	fmt.Println(ReverseString("snakes"))
	fmt.Println(ReverseString("big_angry_fluffy_lions_are_eating_raw_meat"))
	fmt.Println(ReverseString("главрыба"))
	fmt.Println(ReverseString("狐犬"))
}
