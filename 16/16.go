//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package main

import "fmt"

// Начало выполнения алгоритма, в котором вызывается quickSort на всю длину исходного массива.
func quickStart(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		//Массив разделяется на две половины, и для каждой потом выполняется сортировка с продолжением разбиения.
		p := partition(arr, low, high)
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}
func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{1, 45, 2, 3, -67, 123, 23, 14, 78, -4}
	quickStart(arr)
	fmt.Println("Sorted array 1: ", arr)

	arr2 := []int{1, 1, 23, -17, 3, -78, -43, 12, 15, 18, -2345678}
	quickStart(arr2)
	fmt.Println("Sorted array 2: ", arr2)
}
