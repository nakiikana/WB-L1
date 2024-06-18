// Способ остановки горутины 2 (канал для работы и сигнализации о прекращении работы)

package main

import (
	"fmt"
	"sync"
	"time"
)

// Один и тот же канал можно использовать
// и для передачи данных, и для сигнализации о завершении работы.
func worker(ch chan bool, wg *sync.WaitGroup) {
	//Заводим отдельную горутину, которая будет отслеживать, закрыт ли канал. В случае закрытия
	//прекращать всю работу.
	go func() {
		for {
			_, ok := <-ch
			if !ok {
				fmt.Println("Goroutine is finishing its work")
				wg.Done()
				return
			}
		}
	}()
	//И еще одну, выполняющую работу.
	go func() {
		for {
			fmt.Println("Goroutine is doing its work")
			time.Sleep(1 * time.Second)
		}
	}()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan bool)

	go worker(ch, &wg)

	time.Sleep(5 * time.Second)
	close(ch)
	wg.Wait()
}
