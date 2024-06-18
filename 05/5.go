// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeLimitTask(ctx context.Context, ch chan interface{}, wg *sync.WaitGroup) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				log.Println("Finished work with timeout")
				close(ch)
				return
			default:
				wg.Add(2)
				go WriteToChannel(ch, wg)
				go ReadFromChannel(ch, wg)
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
func WriteToChannel(ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- rand.Intn(100)
}

func ReadFromChannel(ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	val := <-ch
	log.Println("Read a new value from channel: ", val)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan interface{})
	wg.Add(1)
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Could not read the amount of seconds: %v\n", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()
	timeLimitTask(ctx, ch, &wg)
	wg.Wait()
}
