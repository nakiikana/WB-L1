// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться.
package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/tjgq/ticker"
)

func readAndWrite(n int, ch chan interface{}) {

}

func GracefulSlaughter(ch chan interface{}, n int, wg sync.WaitGroup) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.)
	tick := ticker.NewTicker(time.Duration(n) * time.Second)
	go func() {
		for {
			select {
			case <-tick:
				close(ch)
				wg.Wait()
			}
			case 
		}
	}()
}

func main() {
	ch := make(chan interface{})
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Could not read the amount of seconds: %v\n", err)
	}
	readAndWrite(n, ch)
}
