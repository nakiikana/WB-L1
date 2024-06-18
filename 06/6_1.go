// Способ остановки горутины 1 (отдельный канал)

package main

import (
	"fmt"
	"sync"
	"time"
)

// Главным способом коммуникации горутин (в т.ч. для прекращения работы) являются каналы.
// В данном примере создается и используется отдельный канал для сигнализирования о выходе из цикла.
func worker(ch chan bool, wg *sync.WaitGroup) {
	for {
		select {
		// Как только в канал поступает bool, работа горутины прекращается.
		case <-ch:
			wg.Done()
			close(ch)
			fmt.Println("Goroutine is finishing its work")
			return
		default:
			//В противном случае выполняем необходимые действия.
			fmt.Println("Some work is being done")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan bool)
	wg.Add(1)
	go worker(ch, &wg)

	time.Sleep(10 * time.Second)
	ch <- true
}
