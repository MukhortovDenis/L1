package main

import (
	"fmt"
	"sync"
)

func main() {
	var counters = map[int]int{}
	mu := &sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(counters map[int]int, i int, mu *sync.Mutex) {
			for j := 0; j < 5; j++ {
				mu.Lock()
				counters[i*10+j] = j
				mu.Unlock()
			}
			wg.Done()
		}(counters, i, mu)
	}
	wg.Wait()
	fmt.Println("Результат", counters)
}
