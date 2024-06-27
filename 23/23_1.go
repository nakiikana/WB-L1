//Удалить i-ый элемент из слайса

package main

import "fmt"

func DeleteFromSlice(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 5, 12, 111, 17, 19, 25, 29, 32}
	removed := DeleteFromSlice(slice, 3)
	fmt.Println("The resulting array after romoval of the third element: ", removed)
	//Из исходного массива элемент также пропадает, но сохраняется его длина, а последний элемент дублируется:
	fmt.Println("The initial array after romoval of the third element: ", slice)
	//Кроме того, если мы поменяем один из элементов нового слайса, то изменится и исходный, так как низлежащий массив один.
	removed[0] = 0
	fmt.Println("The new one: ", removed)
	fmt.Println("The initial one: ", slice)
	//При сохранении изменений в исходный массив проблема с дублированием последнего элемента решается
	slice = DeleteFromSlice(slice, 2)
	fmt.Println("Change saved to the initial slice: ", slice)
	//Однако если мы хотим избежать зависимости между первоначальным слайсом и измененным, то
	//необходимо создать новый слайс и работать с другой областью памяти (23_2.go).
}
