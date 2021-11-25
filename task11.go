package main

import "fmt"

func main() {
	a1 := []int{2, 6, 13, 5, 14, 7, 8, 1, 4, 12}
	a2 := []int{1, 9, 3, 15, 0, 2, 8}
	a3 := newSlice(a1, a2)
	fmt.Println(a3)
}

func newSlice(a1 []int, a2 []int) []int {
	a3 := make([]int, 0, 3)
	for i := range a1 {
		for j := range a2 {
			if a1[i] == a2[j] {
				a3 = append(a3, a1[i])
			}
		}
	}
	return a3
}
