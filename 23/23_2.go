package main

import "fmt"

// Функция для удаления элемента из слайса
func RemoveFromSlice(slice []int, pos int) []int {
	//Выделяем новую область памяти под слайс.
	removed := make([]int, 0, len(slice)-2)
	//Сохраняем в него все элементы, кроме того, что имеет индекс pos.
	for i := range slice {
		if i == pos {
			continue
		}
		removed = append(removed, slice[i])
	}
	return removed
}
func main() {
	sliceNew := []int{1, 2, 3, 4, 5}
	removedNew := RemoveFromSlice(sliceNew, 3)
	fmt.Println("The new initial slice: ", sliceNew)
	fmt.Println("The new slice with a removed element: ", removedNew)
}
