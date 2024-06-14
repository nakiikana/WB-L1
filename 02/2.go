//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func squareArray(arr []int) {
	var wg sync.WaitGroup

	vals := make(chan int, len(arr))

	wg.Add(len(arr))
	for _, v := range arr {
		go func(val int) {
			defer wg.Done()
			vals <- val * val
		}(v)
	}

	wg.Wait()
	close(vals)

	for squareVal := range vals {
		fmt.Println(squareVal)
	}
}

func main() {
	vals := []int{2, 4, 6, 8, 10}
	squareArray(vals)
}
