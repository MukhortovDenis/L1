package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mas := [10]int{1, 5, 14, 7, 23, 1234, 6, 7, 24, 65}
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	for i := range mas {
		ch1 <- mas[i]
	}
	wg.Add(2)
	go func() {
		var x int
		var i int = 0
		for {
			select {
			case x = <-ch1:
				ch2 <- x
				i++
			default:
			}
			if i == 10 {
				close(ch1)
				wg.Done()
				return
			}
		}
	}()
	go func() {
		var i int = 0
		for {
			select {
			case y := <-ch2:
				fmt.Fprint(os.Stdout, y)
				i++
			default:
			}
			if i == 10 {
				close(ch2)
				wg.Done()
				return
			}

		}

	}()
	wg.Wait()
}
