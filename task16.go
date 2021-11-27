package main

import (
	"fmt"
)

func main() {
	a := []int{23, 52, 154, 1, 65, 76, 89, 34, 22}
	fmt.Print(quicksort(a))
}

func quicksort(a []int) []int{
	//Если в слайсе меньше двух элементов, то ничего не делаем
	//Так же с помощью этой проверки останавливаем рекурсию
	if len(a) < 2 {
		return a
	}
	var n int
	l, r := 0, len(a)-1
	if (l+r)%2 == 0 {
		n = (l + r) / 2
	} else {
		n = (l + r + 1) / 2
	}
	a[n], a[r] = a[r], a[n]
	for i := range a{
		if a[i] < a[r]{
			a[l], a[i] = a[i], a[l]
			l++
		}
	}
	a[l], a[r] = a[r], a[l]
	//рекурсивно выполняем сортировку для кусков
	quicksort(a[:l])
	quicksort(a[l+1:])

	return a
}
