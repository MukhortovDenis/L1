package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var duration time.Duration = 20 * time.Second
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	ch := make(chan int, 1)
	defer close(ch)
	go func() {
		var a int
	Loop1:
		for {
			select {
			case a = <-ch:
				fmt.Println(a)
			case <-ctx.Done():
				fmt.Println("Cчитывание закончилось")
				break Loop1
				
			}
		}
	}()
Loop2:
	for i := 0; ; i++ {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			break Loop2

		case ch <- i:
		}
	}
}
