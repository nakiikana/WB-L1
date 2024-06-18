// Способ остановки горутины 4 (использовние пакета tomb)

package main

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/tomb.v1"
)

// Можно остановить работу горутины с помощью пакета tomb, который позволяет отслеживать
// жизненный цикл горутины (она может быть живой, умирающей и мертвой) и причину ее смерти.
func worker(tomb *tomb.Tomb) {
	defer tomb.Done()
	for {
		select {
		case <-tomb.Dying():
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Goroutine is doing its work")
		}
	}
}

func main() {
	testTomb := tomb.Tomb{}

	go worker(&testTomb)

	time.Sleep(10 * time.Second)
	//После вызова Kill новые сообщения перестают появляться.
	testTomb.Kill(errors.New("finish of goroutine work"))
	err := testTomb.Wait()
	fmt.Println(err)
}
