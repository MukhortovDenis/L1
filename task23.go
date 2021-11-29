package main

import (
	"fmt"
	"os"
)

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var i int
	fmt.Fscan(os.Stdin, &i)
	a[i] = a[len(a)- 1]
	a = a[:len(a)-1]
	fmt.Print(a)
}
