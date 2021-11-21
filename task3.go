package main

import (
	"fmt"
	"sync"
)

func square(mas [5]int, i int) int {
	b := mas[i] * mas[i]
	return b
}
func main() {
	var wg sync.WaitGroup
	mas := [5]int{2, 4, 6, 8, 10}
	// mas1 := [5]int{}
	ch := make(chan int, 5)
	for i := 0; i <= 4; i++ {
		go func(i int) {
			wg.Add(1)
			ch <- square(mas, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	answer := <-ch + <-ch + <-ch + <-ch + <-ch
	fmt.Println(answer)
}
