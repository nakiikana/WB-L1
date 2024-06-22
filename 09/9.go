//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
// из массива, во второй — результат операции x*2, после чего данные
// из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
	"sync"
)

type NumberConveyor struct {
	in  chan int
	out chan int
	wg  sync.WaitGroup
}

func NewNumberConveyor(arr []int) *NumberConveyor {
	return &NumberConveyor{
		in:  make(chan int, len(arr)),
		out: make(chan int, len(arr)),
		wg:  sync.WaitGroup{},
	}
}

func (c *NumberConveyor) Write(arr []int) {
	go func() {
		for _, v := range arr {
			c.in <- v
		}
		close(c.in)
	}()
}

func (c *NumberConveyor) Multiply() {
	go func() {
		for v := range c.in {
			c.out <- v * 2
		}
		close(c.out)
	}()
}

func main() {
	var wg sync.WaitGroup

	arr := []int{1, 2, 3, 4, 5}
	c := NewNumberConveyor(arr)
	wg.Add(2)
	go func() {
		defer wg.Done()
		c.Write(arr)
	}()
	go func() {
		defer wg.Done()
		c.Multiply()
	}()
	for v := range c.out {
		fmt.Println(v)
	}
	wg.Wait()
}
