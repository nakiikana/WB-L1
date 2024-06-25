// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Counter struct {
	c  int
	mu *sync.Mutex
	wg *sync.WaitGroup
}

func NewCounter(wg *sync.WaitGroup) *Counter {
	return &Counter{c: 0, mu: &sync.Mutex{}, wg: wg}
}

// Блокируем доступ к переменной для других горутин для избежания ошибки.
func (c *Counter) Add() {
	defer c.wg.Done()
	c.mu.Lock()
	c.c++
	c.mu.Unlock()
}

// Запускаем работу n рутин.
func (c *Counter) Start(n int) {
	for i := 0; i < n; i++ {
		go func() {
			c.Add()
		}()
	}
}

func (c *Counter) GetCurrentValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func main() {
	var wg sync.WaitGroup

	if len(os.Args) < 2 {
		log.Fatalln("You forgot to enter number of increases :C for councurrent work")
	}
	numOfAdds, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Could not read number of workers")
	}
	wg.Add(numOfAdds)
	c := NewCounter(&wg)
	go c.Start(numOfAdds)
	wg.Wait()
	//Итоговое значение счетчика.
	fmt.Println(c.GetCurrentValue())
}
