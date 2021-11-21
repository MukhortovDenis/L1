package main

import (
	"context"
	"fmt"
	"os"
)

var workersNumber int

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	workersNumber = 10
	ch := make(chan int, 1)
	go func() {
		for i := 0; workersNumber > i; i++ {
			go func(ctx context.Context, i int) {
			LOOP:
				for {
					select {
					case <-ctx.Done():
						fmt.Fprintln(os.Stdout, "Горутина", i, "отработала")
						break LOOP
					default:
						fmt.Fprintln(os.Stdout, <-ch)
					}
				}
			}(ctx, i)
		}
	}()
	defer cancel()
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			break
		default:
			ch <- i
		}
	}
}
