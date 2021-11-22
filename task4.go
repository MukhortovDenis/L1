package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var workersNumber int

func main() {
	var wg sync.WaitGroup
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()
	workersNumber = 10
	ch := make(chan int)
	defer close(ch)
	exitChan := make(chan struct{}, 1)
	defer close(exitChan)
	go func() {
		for i := 0; workersNumber > i; i++ {
			wg.Add(1)
			go func(i int, ext chan struct{}) {
			LOOP:
				for {
					select {
					case <-ctx.Done():
						fmt.Println("Горутина", i, "отработала")
						wg.Done()
						break LOOP
					case a := <-ch:
						fmt.Fprintln(os.Stdout, a)
					}
				}
				ext <- struct{}{}
			}(i, exitChan)
		}
	}()
LOOP1:
	for i := 0; ; i++ {
		select {
		case <-exitChan:
			break LOOP1
		case ch <- i:
		}
	}
	wg.Wait()
	fmt.Println("Программа отработала")
}
