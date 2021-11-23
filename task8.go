package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var i int64
	fmt.Fscan(os.Stdin, &i)
	x := intToBit(i)
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}
	fmt.Println(x)
	fmt.Println("Укажите какой бит и какое значени вы хотите поместить")
	var a, b int64
	fmt.Fscan(os.Stdin, &a, &b)
	if b != 0 && b != 1 {
		fmt.Println("Указали неправильное значение бита")
		return
	}
	l := int64(len(x))
	if a >= l {
		fmt.Println("out of range")
		return
	}
	x[a] = b
	fmt.Println(x)
	var newnumb int64 = 0
	i = bitToInt(x, l, newnumb)
	fmt.Println(i)

}
func intToBit(i int64) []int64 {
	s := make([]int64, 0)
	for {
		if i%2 == 0 && i != 1 && i != 0 {
			s = append(s, i%2)
			i = i / 2
		} else if i%2 != 0 && i != 1 && i != 0 {
			s = append(s, i%2)
			i = (i - i%2) / 2
		} else if i == 1 {
			s = append(s, i)
			i = i - i
		} else {
			return s
		}
	}
}
func bitToInt(x []int64, l int64, newnumb int64) int64 {
	for i := range x {
		newnumb += x[i] * int64(math.Pow(2, float64(l)-1))
		l--
	}
	return newnumb
}
