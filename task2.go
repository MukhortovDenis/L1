package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	mas := [5]int{2, 4, 6, 8, 10}
	for i := 0; i < 5; i++{
		wg.Add(1)
		go func(i int) {
			fmt.Fprintln(os.Stdout, mas[i]*mas[i])
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Fprintln(os.Stdout, "Программа отработала")
}
