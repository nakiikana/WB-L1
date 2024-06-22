// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func squareSum(arr []int) int {
	var wg sync.WaitGroup
	sum := 0
	wg.Add(len(arr))
	for _, v := range arr {
		go func(val int) {
			defer wg.Done()
			sum += val * val
		}(v)
	}
	wg.Wait()
	return sum
}

func main() {
	vals := []int{2, 4, 6, 8, 10}
	res := squareSum(vals)
	fmt.Println(res)
}
