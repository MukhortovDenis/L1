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
	ch := make(chan int, 5)
	for i := 0; i <= 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- square(mas, i)
		}(i)
	}
	wg.Wait()
	answer := <-ch + <-ch + <-ch + <-ch + <-ch
	fmt.Println(answer)
}
