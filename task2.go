package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	mas := [5]int{2, 4, 6, 8, 10}
	go func(){
		fmt.Fprintln(os.Stdout, mas[0]*mas[0])
		wg.Done()
	}()
	go func(){
		fmt.Fprintln(os.Stdout, mas[1]*mas[1])
		wg.Done()
	}()
	go func(){
		fmt.Fprintln(os.Stdout, mas[2]*mas[2])
		wg.Done()
	}()
	go func(){
		fmt.Fprintln(os.Stdout, mas[3]*mas[3])
		wg.Done()
	}()
	go func(){
		fmt.Fprintln(os.Stdout, mas[4]*mas[4])
		wg.Done()
	}()
	wg.Wait()
	fmt.Fprintln(os.Stdout, "Программа отработала")
}
