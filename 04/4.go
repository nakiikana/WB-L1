// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N
// воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	id   int
	task chan interface{}
	wg   *sync.WaitGroup
}

type Pool struct {
	pool []Worker
}

func newPool(max int, task chan interface{}, wg *sync.WaitGroup) Pool {
	pool := make([]Worker, 0, max)
	for i := 0; i < max; i++ {
		pool = append(pool, Worker{id: i, task: task, wg: wg})
	}
	poolWorker := Pool{pool: pool}
	return poolWorker
}

func (p *Pool) Read() {
	var wg sync.WaitGroup
	for _, v := range p.pool {
		go func(w Worker) {
			defer w.wg.Done()
			for d := range w.task {
				fmt.Printf("worker #%d received data: %v\n", w.id, d)
			}
		}(v)
	}
	wg.Wait()
}

func GracefulSlaughter(ch chan interface{}, wg *sync.WaitGroup) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	log.Println("End of work")

	close(ch)
	wg.Wait()
}

func main() {

	var wg sync.WaitGroup
	task := make(chan interface{})
	workerNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Could not read number of workers from the command line: ", err)
	}
	wg.Add(workerNum)

	go func() {
		for {
			task <- RandStringRunes(5)
			task <- rand.Intn(100)
			time.Sleep(2 * time.Second)
		}
	}()
	pool := newPool(workerNum, task, &wg)
	pool.Read()

	GracefulSlaughter(task, &wg)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
