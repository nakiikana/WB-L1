// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N
// воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
package main

import (
	"os"
	"sync"
)

type Worker struct {
	task chan interface{}
}

type Pool struct {
	pool []Worker
}

func newPool(max int, task chan interface{}) Pool {
	pool := make([]Worker, 0, max)
	for i := 0; i < max; i++ {
		pool = append(pool, Worker{task: task})
	}
	poolWorker := Pool{pool: pool}
	return pool
}

func poolWorkerRead(workerCount int, in chan interface{}) {
	var wg sync.WaitGroup
	for i := 0; i <= workerCount; i++ {
		wg.Add(1)
	}
	wg.Wait()
}

func main() {
	workerNum := os.Args[0]
}
