//Реализовать бинарный поиск встроенными методами языка.

package main

func BinarySort(arr []int, n int) int {
	right := 0
	left := len(arr) - 1
	for right < left {
		if arr[left/2] == n {
			return left / 2
		}
		if arr[left/2] > n {
			right
		}
	}
}
