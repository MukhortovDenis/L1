package main

import (
	"fmt"
	"math"
)

func main() {
	mas := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// sort.Sort(sort.Float64Slice(mas))
	for i := -20.0; i < 40; {
		fmt.Println(i, " : ", calc(mas, i))
		i = i + 10
	}
}
func calc(mas []float64, i float64) []float64 {
	a := make([]float64, 0, 3)
	for y := range mas {
		if math.Mod(mas[y], i) < 10 && mas[y]/i > 0 && mas[y]-i < 10 && mas[y]/i < 2 {
			a = append(a, mas[y])
		}

	}
	return a
}
