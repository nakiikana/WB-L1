// Способ остановки горутины 3 (контекст)

package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, ch chan struct{}) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			ch <- struct{}{}
			fmt.Println("Goroutine is finishing its work")
			return
		default:
			fmt.Println("Goroutine is doing its work")
		}
	}
}

func main() {
	ch := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, ch)

	go func() {
		time.Sleep(5 * time.Second)
		cancel()
	}()
	<-ch
	fmt.Println("The end")
}
