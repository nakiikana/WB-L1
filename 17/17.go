//Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

// Бинарный поиск эффективен на отсортированном массиве чисел. В худшем случае сложность - O(log n),
// в лучшем - O(1).
func BinarySearch(arr []int, n int) int {
	//Устанавливаем границы поиска
	right := len(arr) - 1
	left := 0
	for left <= right {
		//Находим значение, расположенное в середине массива.
		middle := int((right + left) / 2)
		//Если данное значение совпадает с искомым, то возвращаем его.
		if arr[middle] == n {
			return middle
		}
		//В противном случае сравнимаем с искомым: если значение по середине больше, то
		//рассматриваем крайнюю левую половину массива, иначе - правую.
		if arr[middle] > n {
			right = middle - 1
			continue
		}
		left = middle + 1
	}
	return -1
}

func main() {
	arr := []int{-1, 6, 89, 123, 213, 1234, 2222}
	fmt.Println("1234 -", BinarySearch(arr, 1234), "position")
	fmt.Println("-1 -", BinarySearch(arr, -1), "position")
	fmt.Println("777 -", BinarySearch(arr, 777), "position (not found)")
}
