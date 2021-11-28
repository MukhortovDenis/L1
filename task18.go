package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var gorNum int

type Counter struct {
	mu    sync.Mutex
	count int
}

func main() {
	wg := sync.WaitGroup{}
	counter := Counter{count: 0}
	gorNum = 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for i := 0; i < gorNum; i++ {
		go func(ctx context.Context) {
			wg.Add(1)
			for {
				select {
				case <-ctx.Done():
					wg.Done()
					return
				default:
					counter.mu.Lock()
					counter.count++
					counter.mu.Unlock()
				}
			}
		}(ctx)
	}
	wg.Wait()
	fmt.Print(counter.count)

}
