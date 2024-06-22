// Реализовать конкурентную запись данных в map.
// В данном решении используется sync.Mutex, который блокирует использование переменной другими горутинами, пока текущая не завершит работу.

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func concurrentWrite(m map[string]int) {
	mut := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			defer wg.Done()
			key := RandStringRunes(4)
			mut.Lock()
			m[key] = i
			mut.Unlock()
		}(i)
	}
	wg.Wait()
}

func main() {
	m := make(map[string]int)
	concurrentWrite(m)
	fmt.Println(m)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
